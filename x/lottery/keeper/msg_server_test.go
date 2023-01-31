package keeper_test

import (
	"context"
	"testing"

	keepertest "lottery/testutil/keeper"
	"lottery/x/lottery"
	"lottery/x/lottery/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (lottery.MsgServer, context.Context) {
	k, ctx := keepertest.LotteryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
