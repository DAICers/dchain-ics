package keeper

import (
	"context"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetWhiteboardPixelColor(goCtx context.Context, msg *types.MsgSetWhiteboardPixelColor) (*types.MsgSetWhiteboardPixelColorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	whiteboard, found := k.GetWhiteboard(ctx, msg.WhiteboardId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "whiteboard %d doesn't exist", msg.WhiteboardId)
	}

	//Check if Whiteboard is locked
	if whiteboard.IsLocked {
		if msg.Creator != whiteboard.Owner {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Whiteboard is locked - only owner is allowed to change pixel state")
		}
	}

	if msg.X >= whiteboard.XSize || msg.Y >= whiteboard.YSize {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Whiteboard pixel position out of range")
	}

	if msg.R > 255 || msg.G > 255 || msg.B > 255 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Enter RGB colors - range 0-255")
	}

	//Calculate Pixel Id
	var row uint64 = msg.X
	var col uint64 = msg.Y

	var i uint64 = row*whiteboard.YSize + col

	//Check if Pixel Exits
	pixelmap, found := k.GetWhiteboardPixelMap(ctx, msg.WhiteboardId, i)
	if !found {
		//Pixel does not exist -> create new one

		//Get Next Pixel id
		nextpixelid := k.GetWhiteboardPixelCount(ctx)

		//Create Pixel
		var pixel = types.WhiteboardPixel{
			WhiteboardId: msg.WhiteboardId,
			Index:        i,
			LastModifier: msg.Creator,
			R:            msg.R,
			G:            msg.G,
			B:            msg.B,
		}

		k.AppendWhiteboardPixel(ctx, pixel)

		//Create Pixelmap
		var new_pixelmap = types.WhiteboardPixelMap{
			WhiteboardId:      msg.WhiteboardId,
			Index:             i,
			WhiteboardPixelId: nextpixelid,
		}

		k.SetWhiteboardPixelMap(ctx, new_pixelmap)

	} else {
		//Pixel Exists -> Update

		pixelid := pixelmap.WhiteboardPixelId

		pixel, foundpixel := k.GetWhiteboardPixel(ctx, pixelid)
		if !foundpixel {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "Whiteboard Pixel id: %d not found", pixelid)
		}

		pixel.R = msg.R
		pixel.G = msg.G
		pixel.B = msg.B
		pixel.LastModifier = msg.Creator

		k.SetWhiteboardPixel(ctx, pixel)
	}

	return &types.MsgSetWhiteboardPixelColorResponse{}, nil
}
