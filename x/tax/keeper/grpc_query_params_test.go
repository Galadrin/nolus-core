package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "gitlab-nomo.credissimo.net/nomo/nolus-core/testutil/keeper"
	"gitlab-nomo.credissimo.net/nomo/nolus-core/x/tax/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.TaxKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

func TestParamsQueryNilRequest(t *testing.T) {
	keeper, ctx := testkeeper.TaxKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, nil)
	require.Error(t, err)
	require.Nil(t, response)
}
