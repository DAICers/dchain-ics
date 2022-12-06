package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/DAICers/dchain-ics/testutil/keeper"
	"github.com/DAICers/dchain-ics/testutil/nullify"
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestWhiteboardPixelMapQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWhiteboardPixelMap(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetWhiteboardPixelMapRequest
		response *types.QueryGetWhiteboardPixelMapResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetWhiteboardPixelMapRequest{
				WhiteboardId: msgs[0].WhiteboardId,
				Index:        msgs[0].Index,
			},
			response: &types.QueryGetWhiteboardPixelMapResponse{WhiteboardPixelMap: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetWhiteboardPixelMapRequest{
				WhiteboardId: msgs[1].WhiteboardId,
				Index:        msgs[1].Index,
			},
			response: &types.QueryGetWhiteboardPixelMapResponse{WhiteboardPixelMap: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetWhiteboardPixelMapRequest{
				WhiteboardId: 100000,
				Index:        100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.WhiteboardPixelMap(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestWhiteboardPixelMapQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWhiteboardPixelMap(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllWhiteboardPixelMapRequest {
		return &types.QueryAllWhiteboardPixelMapRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WhiteboardPixelMapAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WhiteboardPixelMap), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.WhiteboardPixelMap),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WhiteboardPixelMapAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WhiteboardPixelMap), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.WhiteboardPixelMap),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.WhiteboardPixelMapAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.WhiteboardPixelMap),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.WhiteboardPixelMapAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
