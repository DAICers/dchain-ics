package keeper

import (
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MaxXSize(ctx),
		k.MaxYSize(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MaxXSize returns the MaxXSize param
func (k Keeper) MaxXSize(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxXSize, &res)
	return
}

// MaxYSize returns the MaxYSize param
func (k Keeper) MaxYSize(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxYSize, &res)
	return
}
