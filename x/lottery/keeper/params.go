package keeper

import (
	"lottery/x/lottery"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) lottery.Params {
	// var params lottery.Params
	// k.paramstore.GetParamSet(ctx, &params)
	// TODO: persist params
	return lottery.DefaultParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params lottery.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
