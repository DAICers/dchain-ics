package keeper_test

import (
	"testing"

	testkeeper "github.com/DAICers/dchain-ics/testutil/keeper"
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.WhiteboardKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.MaxXSize, k.MaxXSize(ctx))
	require.EqualValues(t, params.MaxYSize, k.MaxYSize(ctx))
}
