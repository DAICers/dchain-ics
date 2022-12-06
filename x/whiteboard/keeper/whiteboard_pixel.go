package keeper

import (
	"encoding/binary"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetWhiteboardPixelCount get the total number of whiteboardPixel
func (k Keeper) GetWhiteboardPixelCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WhiteboardPixelCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetWhiteboardPixelCount set the total number of whiteboardPixel
func (k Keeper) SetWhiteboardPixelCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WhiteboardPixelCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendWhiteboardPixel appends a whiteboardPixel in the store with a new id and update the count
func (k Keeper) AppendWhiteboardPixel(
	ctx sdk.Context,
	whiteboardPixel types.WhiteboardPixel,
) uint64 {
	// Create the whiteboardPixel
	count := k.GetWhiteboardPixelCount(ctx)

	// Set the ID of the appended value
	whiteboardPixel.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardPixelKey))
	appendedValue := k.cdc.MustMarshal(&whiteboardPixel)
	store.Set(GetWhiteboardPixelIDBytes(whiteboardPixel.Id), appendedValue)

	// Update whiteboardPixel count
	k.SetWhiteboardPixelCount(ctx, count+1)

	return count
}

// SetWhiteboardPixel set a specific whiteboardPixel in the store
func (k Keeper) SetWhiteboardPixel(ctx sdk.Context, whiteboardPixel types.WhiteboardPixel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardPixelKey))
	b := k.cdc.MustMarshal(&whiteboardPixel)
	store.Set(GetWhiteboardPixelIDBytes(whiteboardPixel.Id), b)
}

// GetWhiteboardPixel returns a whiteboardPixel from its id
func (k Keeper) GetWhiteboardPixel(ctx sdk.Context, id uint64) (val types.WhiteboardPixel, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardPixelKey))
	b := store.Get(GetWhiteboardPixelIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWhiteboardPixel removes a whiteboardPixel from the store
func (k Keeper) RemoveWhiteboardPixel(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardPixelKey))
	store.Delete(GetWhiteboardPixelIDBytes(id))
}

// GetAllWhiteboardPixel returns all whiteboardPixel
func (k Keeper) GetAllWhiteboardPixel(ctx sdk.Context) (list []types.WhiteboardPixel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteboardPixelKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WhiteboardPixel
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetWhiteboardPixelIDBytes returns the byte representation of the ID
func GetWhiteboardPixelIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetWhiteboardPixelIDFromBytes returns ID in uint64 format from a byte array
func GetWhiteboardPixelIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
