package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Bet struct {
	Amount sdk.Int `json:"amount,omitempty"`
	Fee    sdk.Int `json:"fee,omitempty"`
	Index  uint64  `json:"index,omitempty"`
}

func NewBet(amount, fee sdk.Coin, index uint64) Bet {
	return Bet{
		Amount: amount.Amount,
		Fee:    fee.Amount,
		Index:  index,
	}
}
