package module_test

import (
	"testing"

	keepertest "lottery/testutil/keeper"
	"lottery/testutil/nullify"
	"lottery/x/lottery"
	"lottery/x/lottery/module"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := lottery.GenesisState{
		Params: lottery.DefaultParams(),
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	module.InitGenesis(ctx, *k, genesisState)
	got := module.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
