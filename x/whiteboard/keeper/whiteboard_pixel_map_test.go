package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/DAICers/dchain-ics/testutil/keeper"
	"github.com/DAICers/dchain-ics/testutil/nullify"
	"github.com/DAICers/dchain-ics/x/whiteboard/keeper"
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWhiteboardPixelMap(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WhiteboardPixelMap {
	items := make([]types.WhiteboardPixelMap, n)
	for i := range items {
		items[i].WhiteboardId = uint64(i)
		items[i].Index = uint64(i)

		keeper.SetWhiteboardPixelMap(ctx, items[i])
	}
	return items
}

func TestWhiteboardPixelMapGet(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboardPixelMap(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWhiteboardPixelMap(ctx,
			item.WhiteboardId,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWhiteboardPixelMapRemove(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboardPixelMap(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhiteboardPixelMap(ctx,
			item.WhiteboardId,
			item.Index,
		)
		_, found := keeper.GetWhiteboardPixelMap(ctx,
			item.WhiteboardId,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestWhiteboardPixelMapGetAll(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboardPixelMap(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWhiteboardPixelMap(ctx)),
	)
}
