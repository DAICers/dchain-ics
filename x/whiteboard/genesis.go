package whiteboard

import (
	"github.com/DAICers/dchain-ics/x/whiteboard/keeper"
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the whiteboardPixel
	for _, elem := range genState.WhiteboardPixelList {
		k.SetWhiteboardPixel(ctx, elem)
	}

	// Set whiteboardPixel count
	k.SetWhiteboardPixelCount(ctx, genState.WhiteboardPixelCount)
	// Set all the whiteboard
	for _, elem := range genState.WhiteboardList {
		k.SetWhiteboard(ctx, elem)
	}

	// Set whiteboard count
	k.SetWhiteboardCount(ctx, genState.WhiteboardCount)
	// Set all the whiteboardPixelMap
	for _, elem := range genState.WhiteboardPixelMapList {
		k.SetWhiteboardPixelMap(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WhiteboardPixelList = k.GetAllWhiteboardPixel(ctx)
	genesis.WhiteboardPixelCount = k.GetWhiteboardPixelCount(ctx)
	genesis.WhiteboardList = k.GetAllWhiteboard(ctx)
	genesis.WhiteboardCount = k.GetWhiteboardCount(ctx)
	genesis.WhiteboardPixelMapList = k.GetAllWhiteboardPixelMap(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
