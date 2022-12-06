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

func (k Keeper) WhiteboardPixelAll(c context.Context, req *types.QueryAllWhiteboardPixelRequest) (*types.QueryAllWhiteboardPixelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var whiteboardPixels []types.WhiteboardPixel
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	whiteboardPixelStore := prefix.NewStore(store, types.KeyPrefix(types.WhiteboardPixelKey))

	pageRes, err := query.Paginate(whiteboardPixelStore, req.Pagination, func(key []byte, value []byte) error {
		var whiteboardPixel types.WhiteboardPixel
		if err := k.cdc.Unmarshal(value, &whiteboardPixel); err != nil {
			return err
		}

		whiteboardPixels = append(whiteboardPixels, whiteboardPixel)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWhiteboardPixelResponse{WhiteboardPixel: whiteboardPixels, Pagination: pageRes}, nil
}

func (k Keeper) WhiteboardPixel(c context.Context, req *types.QueryGetWhiteboardPixelRequest) (*types.QueryGetWhiteboardPixelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	whiteboardPixel, found := k.GetWhiteboardPixel(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetWhiteboardPixelResponse{WhiteboardPixel: whiteboardPixel}, nil
}
