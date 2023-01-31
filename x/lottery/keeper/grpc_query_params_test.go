package keeper_test

import (
	"testing"

	testkeeper "lottery/testutil/keeper"
	"lottery/x/lottery"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := lottery.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &lottery.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &lottery.QueryParamsResponse{Params: params}, response)
}
