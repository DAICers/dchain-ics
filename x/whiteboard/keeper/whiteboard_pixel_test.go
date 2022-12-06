package keeper_test

import (
	"testing"

	keepertest "github.com/DAICers/dchain-ics/testutil/keeper"
	"github.com/DAICers/dchain-ics/testutil/nullify"
	"github.com/DAICers/dchain-ics/x/whiteboard/keeper"
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNWhiteboardPixel(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WhiteboardPixel {
	items := make([]types.WhiteboardPixel, n)
	for i := range items {
		items[i].Id = keeper.AppendWhiteboardPixel(ctx, items[i])
	}
	return items
}

func TestWhiteboardPixelGet(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboardPixel(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetWhiteboardPixel(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestWhiteboardPixelRemove(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboardPixel(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhiteboardPixel(ctx, item.Id)
		_, found := keeper.GetWhiteboardPixel(ctx, item.Id)
		require.False(t, found)
	}
}

func TestWhiteboardPixelGetAll(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboardPixel(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWhiteboardPixel(ctx)),
	)
}

func TestWhiteboardPixelCount(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboardPixel(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetWhiteboardPixelCount(ctx))
}
