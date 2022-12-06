package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLockWhiteboard = "lock_whiteboard"

var _ sdk.Msg = &MsgLockWhiteboard{}

func NewMsgLockWhiteboard(creator string, whiteboardId uint64) *MsgLockWhiteboard {
	return &MsgLockWhiteboard{
		Creator:      creator,
		WhiteboardId: whiteboardId,
	}
}

func (msg *MsgLockWhiteboard) Route() string {
	return RouterKey
}

func (msg *MsgLockWhiteboard) Type() string {
	return TypeMsgLockWhiteboard
}

func (msg *MsgLockWhiteboard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLockWhiteboard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLockWhiteboard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
