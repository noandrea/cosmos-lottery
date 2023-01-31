package keeper

import (
	"lottery/x/lottery"
)

var _ lottery.QueryServer = Keeper{}
