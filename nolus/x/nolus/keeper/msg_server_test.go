package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "gitlab-nomo.credissimo.net/momo/cosmzone/nolus/testutil/keeper"
	"gitlab-nomo.credissimo.net/momo/cosmzone/nolus/x/nolus/keeper"
	"gitlab-nomo.credissimo.net/momo/cosmzone/nolus/x/nolus/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NolusKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
