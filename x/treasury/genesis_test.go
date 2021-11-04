package treasury_test

import (
	"testing"

	keepertest "gitlab-nomo.credissimo.net/nomo/cosmzone/testutil/keeper"
	"gitlab-nomo.credissimo.net/nomo/cosmzone/x/treasury"
	"gitlab-nomo.credissimo.net/nomo/cosmzone/x/treasury/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TreasuryKeeper(t)
	treasury.InitGenesis(ctx, *k, genesisState)
	got := treasury.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
