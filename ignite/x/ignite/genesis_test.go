package ignite_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gitlab-nomo.credissimo.net/momo/cosmzone/ignite/testutil/keeper"
	"gitlab-nomo.credissimo.net/momo/cosmzone/ignite/testutil/nullify"
	"gitlab-nomo.credissimo.net/momo/cosmzone/ignite/x/ignite"
	"gitlab-nomo.credissimo.net/momo/cosmzone/ignite/x/ignite/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgniteKeeper(t)
	ignite.InitGenesis(ctx, *k, genesisState)
	got := ignite.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
