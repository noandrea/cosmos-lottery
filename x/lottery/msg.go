package lottery

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ sdk.Msg = &MsgPlaceBetRequest{}

// NewMsgPlaceBetRequest creates a new MsgPlaceBetRequest instance
func NewMsgPlaceBetRequest(
	amount sdk.Coin,
	signerAccount sdk.AccAddress,
) *MsgPlaceBetRequest {
	return &MsgPlaceBetRequest{
		Amount: amount,
		Signer: signerAccount.String(),
	}
}

// Route implements sdk.Msg
func (MsgPlaceBetRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgPlaceBetRequest) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgPlaceBetRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgPlaceBetRequest) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

func (m *MsgPlaceBetRequest) ValidateBasic() error {
	return nil
}
