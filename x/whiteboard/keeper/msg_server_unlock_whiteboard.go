package keeper

import (
	"context"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UnlockWhiteboard(goCtx context.Context, msg *types.MsgUnlockWhiteboard) (*types.MsgUnlockWhiteboardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	whiteboard, found := k.GetWhiteboard(ctx, msg.WhiteboardId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "Whiteboard %d doesn't exist", msg.WhiteboardId)
	}

	//Check if request is from the whiteboard owner
	if msg.Creator != whiteboard.Owner {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only the whiteboard owner is able to change the lock state")
	}

	if !whiteboard.IsLocked {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Whiteboard %d is already unlocked", msg.WhiteboardId)
	}

	whiteboard.IsLocked = false

	k.SetWhiteboard(
		ctx,
		whiteboard,
	)

	return &types.MsgUnlockWhiteboardResponse{}, nil
}
