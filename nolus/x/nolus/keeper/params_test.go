package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "gitlab-nomo.credissimo.net/momo/cosmzone/nolus/testutil/keeper"
	"gitlab-nomo.credissimo.net/momo/cosmzone/nolus/x/nolus/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NolusKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
