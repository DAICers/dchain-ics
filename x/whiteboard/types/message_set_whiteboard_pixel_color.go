package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetWhiteboardPixelColor = "set_whiteboard_pixel_color"

var _ sdk.Msg = &MsgSetWhiteboardPixelColor{}

func NewMsgSetWhiteboardPixelColor(creator string, whiteboardId uint64, x uint64, y uint64, r uint64, g uint64, b uint64) *MsgSetWhiteboardPixelColor {
	return &MsgSetWhiteboardPixelColor{
		Creator:      creator,
		WhiteboardId: whiteboardId,
		X:            x,
		Y:            y,
		R:            r,
		G:            g,
		B:            b,
	}
}

func (msg *MsgSetWhiteboardPixelColor) Route() string {
	return RouterKey
}

func (msg *MsgSetWhiteboardPixelColor) Type() string {
	return TypeMsgSetWhiteboardPixelColor
}

func (msg *MsgSetWhiteboardPixelColor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetWhiteboardPixelColor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetWhiteboardPixelColor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
