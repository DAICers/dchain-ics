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

func createNWhiteboard(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Whiteboard {
	items := make([]types.Whiteboard, n)
	for i := range items {
		items[i].Id = keeper.AppendWhiteboard(ctx, items[i])
	}
	return items
}

func TestWhiteboardGet(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboard(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetWhiteboard(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestWhiteboardRemove(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboard(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhiteboard(ctx, item.Id)
		_, found := keeper.GetWhiteboard(ctx, item.Id)
		require.False(t, found)
	}
}

func TestWhiteboardGetAll(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboard(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWhiteboard(ctx)),
	)
}

func TestWhiteboardCount(t *testing.T) {
	keeper, ctx := keepertest.WhiteboardKeeper(t)
	items := createNWhiteboard(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetWhiteboardCount(ctx))
}
