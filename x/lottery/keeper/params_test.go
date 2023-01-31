package keeper_test

import (
	"testing"

	testkeeper "lottery/testutil/keeper"
	"lottery/x/lottery"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LotteryKeeper(t)
	params := lottery.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
