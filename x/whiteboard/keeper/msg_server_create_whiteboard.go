package keeper

import (
	"context"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateWhiteboard(goCtx context.Context, msg *types.MsgCreateWhiteboard) (*types.MsgCreateWhiteboardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.XSize >= types.DefaultMaxXSize {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "whiteboard x size to large - max: %d", types.DefaultMaxXSize)
	}

	if msg.XSize >= types.DefaultMaxYSize {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "whiteboard y size to large - max: %d", types.DefaultMaxYSize)
	}

	var whiteboard = types.Whiteboard{
		Owner:    msg.Creator,
		Name:     msg.Name,
		XSize:    msg.XSize,
		YSize:    msg.YSize,
		IsLocked: false,
	}

	k.AppendWhiteboard(
		ctx,
		whiteboard,
	)

	return &types.MsgCreateWhiteboardResponse{}, nil
}
