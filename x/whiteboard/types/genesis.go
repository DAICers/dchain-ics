package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		WhiteboardPixelList:    []WhiteboardPixel{},
		WhiteboardList:         []Whiteboard{},
		WhiteboardPixelMapList: []WhiteboardPixelMap{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in whiteboardPixel
	whiteboardPixelIdMap := make(map[uint64]bool)
	whiteboardPixelCount := gs.GetWhiteboardPixelCount()
	for _, elem := range gs.WhiteboardPixelList {
		if _, ok := whiteboardPixelIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for whiteboardPixel")
		}
		if elem.Id >= whiteboardPixelCount {
			return fmt.Errorf("whiteboardPixel id should be lower or equal than the last id")
		}
		whiteboardPixelIdMap[elem.Id] = true
	}
	// Check for duplicated ID in whiteboard
	whiteboardIdMap := make(map[uint64]bool)
	whiteboardCount := gs.GetWhiteboardCount()
	for _, elem := range gs.WhiteboardList {
		if _, ok := whiteboardIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for whiteboard")
		}
		if elem.Id >= whiteboardCount {
			return fmt.Errorf("whiteboard id should be lower or equal than the last id")
		}
		whiteboardIdMap[elem.Id] = true
	}
	// Check for duplicated index in whiteboardPixelMap
	whiteboardPixelMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.WhiteboardPixelMapList {
		index := string(WhiteboardPixelMapKey(elem.WhiteboardId, elem.Index))
		if _, ok := whiteboardPixelMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for whiteboardPixelMap")
		}
		whiteboardPixelMapIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
