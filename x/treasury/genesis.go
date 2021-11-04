package treasury

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab-nomo.credissimo.net/nomo/cosmzone/x/treasury/keeper"
	"gitlab-nomo.credissimo.net/nomo/cosmzone/x/treasury/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genState := k.GetParams(ctx)
	genesis := types.NewGenesis(genState.FeeRate, genState.FeeCaps, genState.FeeProceeds)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
