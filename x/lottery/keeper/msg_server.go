package keeper

import (
	"context"
	"lottery/x/lottery"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) lottery.MsgServer {
	return &msgServer{
		Keeper: keeper,
	}
}

var _ lottery.MsgServer = msgServer{}

// PlaceBet handle the bid request
func (k msgServer) PlaceBet(
	goCtx context.Context,
	msg *lottery.MsgPlaceBetRequest,
) (*lottery.MsgPlaceBetResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// note, we do validation here instead of validate basic
	// because the validation logic is self contained and easier to test
	// it has the downside that is more expensive for the transaction signer
	// in case of validation errors
	if msg == nil {
		err := sdkerrors.Wrapf(lottery.ErrInvalidBetRequest, "the bet request is missing required values")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	k.Logger(ctx).Info("place bet tx received", "bidder", msg.Signer)
	p := k.GetParams(ctx)

	// verify the bet amount
	if msg.Amount.IsLT(p.MinBetAmount) {
		err := sdkerrors.Wrapf(lottery.ErrLotteryBetTooLow, "min: %s, got: %s", p.MinBetAmount, msg.Amount)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	// verify the bet amount
	if msg.Amount.IsGTE(p.MaxBetAmount.AddAmount(sdk.NewInt(1))) {
		err := sdkerrors.Wrapf(lottery.ErrLotteryBetTooHigh, "max: %s, got: %s", p.MaxBetAmount, msg.Amount)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// parse the account address
	signerAccount, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		err := sdkerrors.Wrapf(lottery.ErrLotteryBetTooHigh, "cannot retrieve the account from address %s", msg.Signer)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// verify the sender balance
	// TODO: use GTE or GT?
	signerBalance := k.GetBalance(ctx, signerAccount, p.LotteryFee.Denom)
	requiredBalance := msg.Amount.Amount.Add(p.LotteryFee.Amount)
	if !signerBalance.GTE(requiredBalance) {
		err := sdkerrors.Wrapf(lottery.ErrInsufficientBalance, "required: %s, got: %s", requiredBalance, signerBalance)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// get the pool size
	poolSize := k.GetPoolSize(ctx)

	// create a new bet
	bet := lottery.NewBet(msg.Amount, p.LotteryFee, poolSize)

	// try to submit the bet
	if betErr := k.Keeper.PlaceBet(ctx, signerAccount, bet, ctx.TxBytes()); betErr != nil {
		err := sdkerrors.Wrapf(lottery.ErrBetFailed, betErr.Error())
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	k.Logger(ctx).Info("place bet tx successful")

	return &lottery.MsgPlaceBetResponse{}, nil
}
