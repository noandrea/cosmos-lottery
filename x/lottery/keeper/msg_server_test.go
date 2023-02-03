package keeper_test

import (
	"context"
	keepertest "lottery/testutil/keeper"
	"lottery/x/lottery"
	"lottery/x/lottery/keeper"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (lottery.MsgServer, context.Context) {
	k, ctx := keepertest.LotteryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func Test_msgServer_PlaceBet(t *testing.T) {

	msgServer, ctx := setupMsgServer(t)

	tests := []struct {
		name    string
		msg     *lottery.MsgPlaceBetRequest
		wantErr error
	}{
		{
			"OK: place-bet successful",
			&lottery.MsgPlaceBetRequest{
				Amount: sdk.NewInt64Coin(sdk.DefaultBondDenom, 10_000_000),
				Signer: "lot1k6g24mg4l5ktpyzt0qd8lmcp27d7u9yj76l6vx",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := msgServer.PlaceBet(ctx, tt.msg)

			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}
