package keeper

import (
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
)

var _ types.QueryServer = Keeper{}
