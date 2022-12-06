package types_test

import (
	"testing"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated whiteboardPixel",
			genState: &types.GenesisState{
				WhiteboardPixelList: []types.WhiteboardPixel{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid whiteboardPixel count",
			genState: &types.GenesisState{
				WhiteboardPixelList: []types.WhiteboardPixel{
					{
						Id: 1,
					},
				},
				WhiteboardPixelCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated whiteboard",
			genState: &types.GenesisState{
				WhiteboardList: []types.Whiteboard{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid whiteboard count",
			genState: &types.GenesisState{
				WhiteboardList: []types.Whiteboard{
					{
						Id: 1,
					},
				},
				WhiteboardCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated whiteboardPixelMap",
			genState: &types.GenesisState{
				WhiteboardPixelMapList: []types.WhiteboardPixelMap{
					{
						WhiteboardId: 0,
						Index:        0,
					},
					{
						WhiteboardId: 0,
						Index:        0,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
