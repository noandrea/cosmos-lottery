package keeper

import (
	"lottery/x/lottery"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) lottery.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ lottery.MsgServer = msgServer{}
