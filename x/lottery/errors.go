package lottery

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lottery module sentinel errors
var (
	ErrInvalidBetRequest   = sdkerrors.Register(ModuleName, 1100, "Malformed bet request")
	ErrBetFailed           = sdkerrors.Register(ModuleName, 1101, "Betting for this lottery has failed")
	ErrLotteryFeeTooLow    = sdkerrors.Register(ModuleName, 1110, "Submitted fee amount is too low")
	ErrLotteryBetTooLow    = sdkerrors.Register(ModuleName, 1120, "Submitted bet amount is too low")
	ErrLotteryBetTooHigh   = sdkerrors.Register(ModuleName, 1130, "Submitted bet amount is too high")
	ErrInsufficientBalance = sdkerrors.Register(ModuleName, 1140, "Account balance insufficient to cover bet costs")
)
