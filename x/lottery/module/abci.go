package module

import (
	"encoding/binary"
	"fmt"
	"time"

	"lottery/x/lottery"
	"lottery/x/lottery/keeper"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	if ctx.BlockHeight() < 1 {
		// the inflation is already calculated for block 1. The inflation of the block 1 esist in the time
		// between block 0 (chain start) and block 1.
		return
	}
	defer telemetry.ModuleMeasureSince(lottery.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// verify that the block proposer is not participating
	if k.HasBet(ctx, string(ctx.BlockHeader().ProposerAddress)) {
		k.Logger(ctx).Info("lottery not fired: proposer is part of the lottery")
		return
	}
	// verify that the pool is big enough
	poolSize := k.GetPoolSize(ctx)
	if poolSize < k.GetParams(ctx).LotteryMinPoolSize {
		k.Logger(ctx).Info("lottery not fired: lottery pool too small", "required", k.GetParams(ctx).LotteryMinPoolSize, "actual", poolSize)
		return
	}
	// get the lottery data
	lh, maxBet, minBet := k.GetLotteryData(ctx)
	// get the first 2 bytes
	v := binary.BigEndian.Uint16(lh[0:2])
	// get the index of the winner
	winnerIndex := uint64(v) % poolSize
	// retrieve the bet and the winner address
	winnerAddress, winnerBet, err := k.GetLotteryWinner(ctx, winnerIndex)
	if err != nil {
		panic(fmt.Errorf("unrecoverable error: %s", err))
	}
	// calculate the payout
	if winnerBet.Amount.Equal(minBet) {
		k.Logger(ctx).Info("lottery fired: winner has min bet - no payout", "winner", winnerAddress.String(), "bet", winnerBet.Amount.String())
		return
	}
	if winnerBet.Amount.Equal(maxBet) {
		k.Logger(ctx).Info("lottery fired: winner has max bet - full payout", "winner", winnerAddress.String(), "bet", winnerBet.Amount.String())
		k.ExecutePayout(ctx, winnerAddress, keeper.PayoutFull)
		return
	}
	k.Logger(ctx).Info("lottery fired: payouts only for bets", "winner", winnerAddress.String(), "bet", winnerBet.Amount.String())
	k.ExecutePayout(ctx, winnerAddress, keeper.PayoutBet)
}
