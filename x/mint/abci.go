package mint

import (
	"errors"
	"fmt"
	"time"

	"github.com/Nolus-Protocol/nolus-core/x/mint/keeper"
	"github.com/Nolus-Protocol/nolus-core/x/mint/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	normInitialTotal     = types.CalcTokensByIntegral(types.NormOffset)
	nanoSecondsInMonth   = sdk.NewDec(time.Hour.Nanoseconds() * 24 * 30)
	nanoSecondsInFormula = types.MonthsInFormula.Mul(nanoSecondsInMonth)
	twelveMonths         = sdk.MustNewDecFromStr("12.0")

	errTimeInFutureBeforeTimePassed = errors.New("time in future can not be before passed time")
)

func calcFunctionIncrement(nanoSecondsPassed sdk.Uint) sdk.Dec {
	return types.NormMonthsRange.Mul(calcFixedIncrement(nanoSecondsPassed))
}

func calcFixedIncrement(nanoSecondsPassed sdk.Uint) sdk.Dec {
	return types.DecFromUint(nanoSecondsPassed).Quo(nanoSecondsInMonth)
}

func calcTimeDifference(blockTime sdk.Uint, prevBlockTime sdk.Uint, maxMintableSeconds sdk.Uint) sdk.Uint {
	if prevBlockTime.GT(blockTime) {
		panic("new block time cannot be smaller than previous block time")
	}

	nsecBetweenBlocks := blockTime.Sub(prevBlockTime)
	if nsecBetweenBlocks.GT(maxMintableSeconds) {
		nsecBetweenBlocks = maxMintableSeconds
	}

	return nsecBetweenBlocks
}

func calcTokens(blockTime sdk.Uint, minter *types.Minter, maxMintableSeconds sdk.Uint) sdk.Uint {
	if minter.TotalMinted.GTE(types.MintingCap) {
		return sdk.ZeroUint()
	}

	if minter.PrevBlockTimestamp.IsZero() {
		// we do not know how much time has passed since the previous block, thus nothing will be mined
		minter.PrevBlockTimestamp = blockTime
		return sdk.ZeroUint()
	}

	nsecPassed := calcTimeDifference(blockTime, minter.PrevBlockTimestamp, maxMintableSeconds)
	if minter.NormTimePassed.LT(types.MonthsInFormula) {
		// First 96 months follow the minting formula
		// As the integral starts from NormOffset (ie > 0), previous total needs to be incremented by predetermined amount
		previousTotal := minter.TotalMinted.Add(normInitialTotal)
		newNormTime := minter.NormTimePassed.Add(calcFunctionIncrement(nsecPassed))
		nextTotal := types.CalcTokensByIntegral(newNormTime)

		delta := nextTotal.Sub(previousTotal)

		return updateMinter(minter, blockTime, newNormTime, delta)
	} else {
		// After reaching 96 normalized time, mint fixed amount of tokens per month until we reach the minting cap
		normIncrement := calcFixedIncrement(nsecPassed)
		delta := sdk.NewUint((normIncrement.Mul(types.DecFromUint(types.FixedMintedAmount))).TruncateInt().Uint64())

		if minter.TotalMinted.Add(delta).GT(types.MintingCap) {
			// Trim off excess tokens if the cap is reached
			delta = types.MintingCap.Sub(minter.TotalMinted)
		}

		return updateMinter(minter, blockTime, minter.NormTimePassed.Add(normIncrement), delta)
	}
}

func updateMinter(minter *types.Minter, blockTime sdk.Uint, newNormTime sdk.Dec, newlyMinted sdk.Uint) sdk.Uint {
	if newlyMinted.LT(sdk.ZeroUint()) {
		// Sanity check, should not happen. However, if this were to happen,
		// do not update the minter state (primary the previous block timestamp)
		// and wait for a new block which should increase the minted amount
		return sdk.ZeroUint()
	}
	minter.NormTimePassed = newNormTime
	minter.PrevBlockTimestamp = blockTime
	minter.TotalMinted = minter.TotalMinted.Add(newlyMinted)
	return newlyMinted
}

// Returns the amount of tokens that should be minted by the integral formula
// for the period between normTimePassed and the timeInFuture.
func predictMintedByIntegral(totalMinted sdk.Uint, normTimePassed, timeAhead sdk.Dec) (sdk.Uint, error) {
	timeAheadNs := timeAhead.Mul(nanoSecondsInMonth).TruncateInt()
	normTimeInFuture := normTimePassed.Add(calcFunctionIncrement(sdk.Uint(timeAheadNs)))
	if normTimePassed.GT(normTimeInFuture) {
		return sdk.ZeroUint(), errTimeInFutureBeforeTimePassed
	}

	if normTimePassed.GTE(types.MonthsInFormula) {
		return sdk.ZeroUint(), nil
	}

	// integral minting is caped to the 96th month
	if normTimeInFuture.GT(types.MonthsInFormula) {
		normTimeInFuture = types.MonthsInFormula
	}

	return types.CalcTokensByIntegral(normTimeInFuture).Sub(normInitialTotal).Sub(totalMinted), nil
}

// Returns the amount of tokens that should be minted during the fixed amount period
// for the period between NormTimePassed and the timeInFuture.
func predictMintedByFixedAmount(totalMinted sdk.Uint, normTimePassed, timeAhead sdk.Dec) (sdk.Uint, error) {
	timeAheadNs := timeAhead.Mul(nanoSecondsInMonth).TruncateInt()

	normTimeInFuture := normTimePassed.Add(calcFunctionIncrement(sdk.Uint(timeAheadNs)))
	if normTimePassed.GT(normTimeInFuture) {
		return sdk.ZeroUint(), errTimeInFutureBeforeTimePassed
	}

	normFixedPeriod := normTimeInFuture.Sub(calcFunctionIncrement(sdk.Uint(nanoSecondsInFormula.TruncateInt())))
	if normFixedPeriod.LTE(sdk.ZeroDec()) {
		return sdk.ZeroUint(), nil
	}

	// convert norm time to non norm time
	fixedPeriod := normFixedPeriod.Sub(types.NormOffset).Quo(types.NormMonthsRange)

	newlyMinted := fixedPeriod.MulInt(sdk.Int(types.FixedMintedAmount))
	// Trim off excess tokens if the cap is reached
	if totalMinted.Add(sdk.Uint(newlyMinted.TruncateInt())).GT(types.MintingCap) {
		return types.MintingCap.Sub(totalMinted), nil
	}

	return sdk.Uint(newlyMinted.TruncateInt()), nil
}

// Returns the amount of tokens that should be minted
// between the NormTimePassed and the timeAhead
// timeAhead expects months represented in decimal form.
func predictTotalMinted(totalMinted sdk.Uint, normTimePassed, timeAhead sdk.Dec) sdk.Uint {
	integralAmount, err := predictMintedByIntegral(totalMinted, normTimePassed, timeAhead)
	if err != nil {
		return sdk.ZeroUint()
	}

	fixedAmount, err := predictMintedByFixedAmount(totalMinted, normTimePassed, timeAhead)
	if err != nil {
		return sdk.ZeroUint()
	}

	return fixedAmount.Add(integralAmount)
}

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	minter := k.GetMinter(ctx)
	if minter.TotalMinted.GTE(types.MintingCap) {
		return
	}

	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	params := k.GetParams(ctx)
	blockTime := ctx.BlockTime().UnixNano()
	coinAmount := calcTokens(sdk.NewUint(uint64(blockTime)), &minter, params.MaxMintableNanoseconds)

	minter.AnnualInflation = predictTotalMinted(minter.TotalMinted, minter.NormTimePassed, twelveMonths)

	ctx.Logger().Debug(fmt.Sprintf("miner: %v total, %v norm time, %v minted", minter.TotalMinted.String(), minter.NormTimePassed.String(), coinAmount.String()))

	k.SetMinter(ctx, minter)
	if coinAmount.GT(sdk.ZeroUint()) {
		// mint coins, update supply
		mintedCoins := sdk.NewCoins(sdk.NewCoin(params.MintDenom, sdk.NewIntFromBigInt(coinAmount.BigInt())))

		err := k.MintCoins(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}

		// send the minted coins to the fee collector account
		err = k.AddCollectedFees(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}

		defer telemetry.ModuleSetGauge(types.ModuleName, float32(coinAmount.Uint64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyDenom, params.MintDenom),
			sdk.NewAttribute(sdk.AttributeKeyAmount, coinAmount.String()),
		),
	)
}
