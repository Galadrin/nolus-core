package ante

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NomoMempoolFeeDecorator will check if the transaction's fee is at least as large
// as the local validator's minimum gasFee (defined in validator config).
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true
// If fee is high enough or not CheckTx, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use NomoMempoolFeeDecorator
type NomoMempoolFeeDecorator struct {
	tk TreasuryKeeper
}

func NewMempoolFeeDecorator(tk TreasuryKeeper) NomoMempoolFeeDecorator {
	return NomoMempoolFeeDecorator{
		tk: tk,
	}
}

func (mfd NomoMempoolFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	gas := feeTx.GetGas()

	// Ensure that the provided fees meet a minimum threshold for the validator,
	// if this is a CheckTx. This is only for local mempool purposes, and thus
	// is only ran on check tx.
	if ctx.IsCheckTx() && !simulate {

		// deduct the nomo fee from feeCoins
		extraFee, feeRemaining, err := ApplyFee(ctx, mfd.tk, feeTx.GetFee())
		if err != nil {
			return ctx, err
		}

		minGasPrices := ctx.MinGasPrices()
		if !minGasPrices.IsZero() {
			requiredFees := make(sdk.Coins, len(minGasPrices))

			// Determine the required fees by multiplying each required minimum gas
			// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
			glDec := sdk.NewDec(int64(gas))
			for i, gp := range minGasPrices {
				fee := gp.Amount.Mul(glDec)
				requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
			}

			// ensure that enough was paid to cover the validator tax after the custom tax was deduced
			if !feeRemaining.IsAnyGTE(requiredFees) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeRemaining.Add(extraFee...), requiredFees.Add(extraFee...))
			}
		}
	}

	return next(ctx, tx, simulate)
}

func ApplyFee(ctx sdk.Context, tk TreasuryKeeper, feeCoins sdk.Coins) (sdk.Coins, sdk.Coins, error) {
	if tk == nil {
		return sdk.Coins{}, feeCoins, nil
	}

	params := tk.GetParams(ctx)
	proceeds := sdk.Coins{}
	if params.FeeRate.IsZero() {
		return proceeds, feeCoins, nil
	}

	// we will deduct the fee from every denomination send
	for _, fee := range feeCoins {
		proceed := sdk.NewCoin(fee.Denom, params.FeeRate.MulInt(fee.Amount).TruncateInt())
		proceeds = proceeds.Add(proceed)
	}

	deductedFees, neg := feeCoins.SafeSub(proceeds)
	if neg {
		return nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeCoins, proceeds)
	}

	return proceeds, deductedFees, nil
}

// NomoDeductFeeDecorator deducts fees from the first signer of the tx
// If the first signer does not have the funds to pay for the fees, return with InsufficientFunds error
// Call next AnteHandler if fees successfully deducted
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator
type NomoDeductFeeDecorator struct {
	ak         AccountKeeper
	tk         TreasuryKeeper
	bankKeeper types.BankKeeper
}

func NewNomoDeductFeeDecorator(ak AccountKeeper, bk types.BankKeeper, tk TreasuryKeeper) NomoDeductFeeDecorator {
	return NomoDeductFeeDecorator{
		ak:         ak,
		tk:         tk,
		bankKeeper: bk,
	}
}

func (dfd NomoDeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.ak.GetModuleAddress(types.FeeCollectorName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.FeeCollectorName))
	}

	feePayer := feeTx.FeePayer()

	deductFeesFrom := feePayer

	deductFeesFromAcc := dfd.ak.GetAccount(ctx, deductFeesFrom)
	if deductFeesFromAcc == nil {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
	}

	// deduct the fees
	if !feeTx.GetFee().IsZero() {
		customFee, reducedFee, err := ApplyFee(ctx, dfd.tk, feeTx.GetFee())
		if err != nil {
			ctx.Logger().Info("Could not deduce custom fees")
		} else {
			dfd.tk.AddProceeds(ctx, customFee)
		}
		err = DeductFees(dfd.bankKeeper, ctx, deductFeesFromAcc, reducedFee)
		if err != nil {
			return ctx, err
		}
	}

	events := sdk.Events{sdk.NewEvent(sdk.EventTypeTx,
		sdk.NewAttribute(sdk.AttributeKeyFee, feeTx.GetFee().String()),
	)}
	ctx.EventManager().EmitEvents(events)

	return next(ctx, tx, simulate)
}

// DeductFees deducts fees from the given account.
func DeductFees(bankKeeper types.BankKeeper, ctx sdk.Context, acc types.AccountI, fees sdk.Coins) error {
	if !fees.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %s", fees)
	}

	err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.FeeCollectorName, fees)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	return nil
}
