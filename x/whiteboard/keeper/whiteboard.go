package keeper

import (
	"encoding/binary"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetWhiteboardCount get the total number of whiteboard
func (k Keeper) GetWhiteboardCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WhiteboardCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetWhiteboardCount set the total number of whiteboard
func (k Keeper) SetWhiteboardCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WhiteboardCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendWhiteboard appends a whiteboard in the store with a new id and update the count
func (k Keeper) AppendWhiteboard(
	ctx sdk.Context,
	whiteboard types.Whiteboard,
) uint64 {
	// Create the whiteboard
	count := k.GetWhiteboardCount(ctx)

	// Set the ID of the appended value
	whiteboard.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardKey))
	appendedValue := k.cdc.MustMarshal(&whiteboard)
	store.Set(GetWhiteboardIDBytes(whiteboard.Id), appendedValue)

	// Update whiteboard count
	k.SetWhiteboardCount(ctx, count+1)

	return count
}

// SetWhiteboard set a specific whiteboard in the store
func (k Keeper) SetWhiteboard(ctx sdk.Context, whiteboard types.Whiteboard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardKey))
	b := k.cdc.MustMarshal(&whiteboard)
	store.Set(GetWhiteboardIDBytes(whiteboard.Id), b)
}

// GetWhiteboard returns a whiteboard from its id
func (k Keeper) GetWhiteboard(ctx sdk.Context, id uint64) (val types.Whiteboard, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardKey))
	b := store.Get(GetWhiteboardIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWhiteboard removes a whiteboard from the store
func (k Keeper) RemoveWhiteboard(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardKey))
	store.Delete(GetWhiteboardIDBytes(id))
}

// GetAllWhiteboard returns all whiteboard
func (k Keeper) GetAllWhiteboard(ctx sdk.Context) (list []types.Whiteboard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Whiteboard
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetWhiteboardIDBytes returns the byte representation of the ID
func GetWhiteboardIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetWhiteboardIDFromBytes returns ID in uint64 format from a byte array
func GetWhiteboardIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
