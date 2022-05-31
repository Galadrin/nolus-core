package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "gitlab-nomo.credissimo.net/momo/cosmzone/ignite/testutil/keeper"
	"gitlab-nomo.credissimo.net/momo/cosmzone/ignite/x/ignite/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IgniteKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
