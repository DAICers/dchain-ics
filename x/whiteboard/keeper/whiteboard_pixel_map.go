package keeper

import (
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetWhiteboardPixelMap set a specific whiteboardPixelMap in the store from its index
func (k Keeper) SetWhiteboardPixelMap(ctx sdk.Context, whiteboardPixelMap types.WhiteboardPixelMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardPixelMapKeyPrefix))
	b := k.cdc.MustMarshal(&whiteboardPixelMap)
	store.Set(types.WhiteboardPixelMapKey(
		whiteboardPixelMap.WhiteboardId,
		whiteboardPixelMap.Index,
	), b)
}

// GetWhiteboardPixelMap returns a whiteboardPixelMap from its index
func (k Keeper) GetWhiteboardPixelMap(
	ctx sdk.Context,
	whiteboardId uint64,
	index uint64,

) (val types.WhiteboardPixelMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardPixelMapKeyPrefix))

	b := store.Get(types.WhiteboardPixelMapKey(
		whiteboardId,
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWhiteboardPixelMap removes a whiteboardPixelMap from the store
func (k Keeper) RemoveWhiteboardPixelMap(
	ctx sdk.Context,
	whiteboardId uint64,
	index uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardPixelMapKeyPrefix))
	store.Delete(types.WhiteboardPixelMapKey(
		whiteboardId,
		index,
	))
}

// GetAllWhiteboardPixelMap returns all whiteboardPixelMap
func (k Keeper) GetAllWhiteboardPixelMap(ctx sdk.Context) (list []types.WhiteboardPixelMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardPixelMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WhiteboardPixelMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
