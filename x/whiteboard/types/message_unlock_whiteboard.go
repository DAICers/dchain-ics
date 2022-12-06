package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnlockWhiteboard = "unlock_whiteboard"

var _ sdk.Msg = &MsgUnlockWhiteboard{}

func NewMsgUnlockWhiteboard(creator string, whiteboardId uint64) *MsgUnlockWhiteboard {
	return &MsgUnlockWhiteboard{
		Creator:      creator,
		WhiteboardId: whiteboardId,
	}
}

func (msg *MsgUnlockWhiteboard) Route() string {
	return RouterKey
}

func (msg *MsgUnlockWhiteboard) Type() string {
	return TypeMsgUnlockWhiteboard
}

func (msg *MsgUnlockWhiteboard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnlockWhiteboard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnlockWhiteboard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
