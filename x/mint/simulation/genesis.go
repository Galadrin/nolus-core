package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"gitlab-nomo.credissimo.net/nomo/nolus-core/x/mint/types"
)

// simulation parameter constants
const (
	MaxMintableNanoseconds = "MaxMintableNanoseconds"
)

// GenMaxMintableNanoseconds generates random MaxMintableNanoseconds in range [1-60)
func GenMaxMintableNanoseconds(r *rand.Rand) uint64 {
	return uint64(time.Second.Nanoseconds() * int64(r.Intn(59)+1))
}

// RandomizedGenState generates a random GenesisState for mint
func RandomizedGenState(simState *module.SimulationState) {
	// minter
	var maxMintableNSecs uint64
	simState.AppParams.GetOrGenerate(
		simState.Cdc, MaxMintableNanoseconds, &maxMintableNSecs, simState.Rand,
		func(r *rand.Rand) { maxMintableNSecs = GenMaxMintableNanoseconds(r) },
	)
	mintDenom := sdk.DefaultBondDenom
	params := types.NewParams(mintDenom, maxMintableNSecs)

	mintGenesis := types.NewGenesisState(types.InitialMinter(), params)

	bz, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}
