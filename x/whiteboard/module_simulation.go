package whiteboard

import (
	"math/rand"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"

	"github.com/DAICers/dchain-ics/testutil/sample"
	whiteboardsimulation "github.com/DAICers/dchain-ics/x/whiteboard/simulation"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = whiteboardsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateWhiteboard = "op_weight_msg_create_whiteboard"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWhiteboard int = 100

	opWeightMsgLockWhiteboard = "op_weight_msg_lock_whiteboard"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLockWhiteboard int = 100

	opWeightMsgUnlockWhiteboard = "op_weight_msg_unlock_whiteboard"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnlockWhiteboard int = 100

	opWeightMsgSetWhiteboardPixelColor = "op_weight_msg_set_whiteboard_pixel_color"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetWhiteboardPixelColor int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	whiteboardGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&whiteboardGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	whiteboardParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMaxXSize), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(whiteboardParams.MaxXSize))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMaxYSize), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(whiteboardParams.MaxYSize))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateWhiteboard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWhiteboard, &weightMsgCreateWhiteboard, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWhiteboard = defaultWeightMsgCreateWhiteboard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWhiteboard,
		whiteboardsimulation.SimulateMsgCreateWhiteboard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLockWhiteboard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLockWhiteboard, &weightMsgLockWhiteboard, nil,
		func(_ *rand.Rand) {
			weightMsgLockWhiteboard = defaultWeightMsgLockWhiteboard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLockWhiteboard,
		whiteboardsimulation.SimulateMsgLockWhiteboard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnlockWhiteboard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnlockWhiteboard, &weightMsgUnlockWhiteboard, nil,
		func(_ *rand.Rand) {
			weightMsgUnlockWhiteboard = defaultWeightMsgUnlockWhiteboard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnlockWhiteboard,
		whiteboardsimulation.SimulateMsgUnlockWhiteboard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetWhiteboardPixelColor int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetWhiteboardPixelColor, &weightMsgSetWhiteboardPixelColor, nil,
		func(_ *rand.Rand) {
			weightMsgSetWhiteboardPixelColor = defaultWeightMsgSetWhiteboardPixelColor
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetWhiteboardPixelColor,
		whiteboardsimulation.SimulateMsgSetWhiteboardPixelColor(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
