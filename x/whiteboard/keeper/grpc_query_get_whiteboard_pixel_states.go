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

func (k Keeper) GetWhiteboardPixelStates(goCtx context.Context, req *types.QueryGetWhiteboardPixelStatesRequest) (*types.QueryGetWhiteboardPixelStatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)

	var pixels []*types.WhiteboardPixel

	pixelStore := prefix.NewStore(store, []byte(types.WhiteboardPixelKey))

	_, err := query.Paginate(pixelStore, req.Pagination, func(key []byte, value []byte) error {
		var pixel types.WhiteboardPixel
		if err := k.cdc.Unmarshal(value, &pixel); err != nil {
			return err
		}

		if pixel.WhiteboardId == req.WhiteboardId {
			pixels = append(pixels, &pixel)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetWhiteboardPixelStatesResponse{Pixels: pixels}, nil
}
