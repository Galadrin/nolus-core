package nolus_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gitlab-nomo.credissimo.net/momo/cosmzone/nolus/testutil/keeper"
	"gitlab-nomo.credissimo.net/momo/cosmzone/nolus/testutil/nullify"
	"gitlab-nomo.credissimo.net/momo/cosmzone/nolus/x/nolus"
	"gitlab-nomo.credissimo.net/momo/cosmzone/nolus/x/nolus/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NolusKeeper(t)
	nolus.InitGenesis(ctx, *k, genesisState)
	got := nolus.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
