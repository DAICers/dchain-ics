package whiteboard_test

import (
	"testing"

	keepertest "github.com/DAICers/dchain-ics/testutil/keeper"
	"github.com/DAICers/dchain-ics/testutil/nullify"
	"github.com/DAICers/dchain-ics/x/whiteboard"
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WhiteboardPixelList: []types.WhiteboardPixel{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		WhiteboardPixelCount: 2,
		WhiteboardList: []types.Whiteboard{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		WhiteboardCount: 2,
		WhiteboardPixelMapList: []types.WhiteboardPixelMap{
			{
				WhiteboardId: 0,
				Index:        0,
			},
			{
				WhiteboardId: 1,
				Index:        1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WhiteboardKeeper(t)
	whiteboard.InitGenesis(ctx, *k, genesisState)
	got := whiteboard.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WhiteboardPixelList, got.WhiteboardPixelList)
	require.Equal(t, genesisState.WhiteboardPixelCount, got.WhiteboardPixelCount)
	require.ElementsMatch(t, genesisState.WhiteboardList, got.WhiteboardList)
	require.Equal(t, genesisState.WhiteboardCount, got.WhiteboardCount)
	require.ElementsMatch(t, genesisState.WhiteboardPixelMapList, got.WhiteboardPixelMapList)
	// this line is used by starport scaffolding # genesis/test/assert
}
