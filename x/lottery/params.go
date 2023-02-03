package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	p := NewParams()

	p.LotteryFee = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(5_000_000))
	p.MinBetAmount = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1_000_000))
	p.MaxBetAmount = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))
	p.LotteryMinPoolSize = 10

	return p
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
