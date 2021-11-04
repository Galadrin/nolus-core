package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "gitlab-nomo.credissimo.net/nomo/cosmzone/x/treasury/types"
    "gitlab-nomo.credissimo.net/nomo/cosmzone/x/treasury/keeper"
    keepertest "gitlab-nomo.credissimo.net/nomo/cosmzone/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TreasuryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
