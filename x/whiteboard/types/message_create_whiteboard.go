package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateWhiteboard = "create_whiteboard"

var _ sdk.Msg = &MsgCreateWhiteboard{}

func NewMsgCreateWhiteboard(creator string, name string, description string, xSize uint64, ySize uint64, isLocked bool) *MsgCreateWhiteboard {
	return &MsgCreateWhiteboard{
		Creator:     creator,
		Name:        name,
		Description: description,
		XSize:       xSize,
		YSize:       ySize,
		IsLocked:    isLocked,
	}
}

func (msg *MsgCreateWhiteboard) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhiteboard) Type() string {
	return TypeMsgCreateWhiteboard
}

func (msg *MsgCreateWhiteboard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhiteboard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhiteboard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
