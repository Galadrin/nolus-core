package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "gitlab-nomo.credissimo.net/momo/cosmzone/ignite/testutil/keeper"
	"gitlab-nomo.credissimo.net/momo/cosmzone/ignite/x/ignite/keeper"
	"gitlab-nomo.credissimo.net/momo/cosmzone/ignite/x/ignite/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IgniteKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
