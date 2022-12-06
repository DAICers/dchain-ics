package keeper

import (
	"context"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhiteboardAll(c context.Context, req *types.QueryAllWhiteboardRequest) (*types.QueryAllWhiteboardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var whiteboards []types.Whiteboard
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	whiteboardStore := prefix.NewStore(store, types.KeyPrefix(types.WhiteboardKey))

	pageRes, err := query.Paginate(whiteboardStore, req.Pagination, func(key []byte, value []byte) error {
		var whiteboard types.Whiteboard
		if err := k.cdc.Unmarshal(value, &whiteboard); err != nil {
			return err
		}

		whiteboards = append(whiteboards, whiteboard)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWhiteboardResponse{Whiteboard: whiteboards, Pagination: pageRes}, nil
}

func (k Keeper) Whiteboard(c context.Context, req *types.QueryGetWhiteboardRequest) (*types.QueryGetWhiteboardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	whiteboard, found := k.GetWhiteboard(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetWhiteboardResponse{Whiteboard: whiteboard}, nil
}
