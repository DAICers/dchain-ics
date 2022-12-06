package keeper

import (
	"context"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhiteboardPixelMapAll(c context.Context, req *types.QueryAllWhiteboardPixelMapRequest) (*types.QueryAllWhiteboardPixelMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var whiteboardPixelMaps []types.WhiteboardPixelMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	whiteboardPixelMapStore := prefix.NewStore(store, types.KeyPrefix(types.WhiteboardPixelMapKeyPrefix))

	pageRes, err := query.Paginate(whiteboardPixelMapStore, req.Pagination, func(key []byte, value []byte) error {
		var whiteboardPixelMap types.WhiteboardPixelMap
		if err := k.cdc.Unmarshal(value, &whiteboardPixelMap); err != nil {
			return err
		}

		whiteboardPixelMaps = append(whiteboardPixelMaps, whiteboardPixelMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWhiteboardPixelMapResponse{WhiteboardPixelMap: whiteboardPixelMaps, Pagination: pageRes}, nil
}

func (k Keeper) WhiteboardPixelMap(c context.Context, req *types.QueryGetWhiteboardPixelMapRequest) (*types.QueryGetWhiteboardPixelMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetWhiteboardPixelMap(
		ctx,
		req.WhiteboardId,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetWhiteboardPixelMapResponse{WhiteboardPixelMap: val}, nil
}
