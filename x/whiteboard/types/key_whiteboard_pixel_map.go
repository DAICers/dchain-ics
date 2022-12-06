package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WhiteboardPixelMapKeyPrefix is the prefix to retrieve all WhiteboardPixelMap
	WhiteboardPixelMapKeyPrefix = "WhiteboardPixelMap/value/"
)

// WhiteboardPixelMapKey returns the store key to retrieve a WhiteboardPixelMap from the index fields
func WhiteboardPixelMapKey(
	whiteboardId uint64,
	index uint64,
) []byte {
	var key []byte

	whiteboardIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(whiteboardIdBytes, whiteboardId)
	key = append(key, whiteboardIdBytes...)
	key = append(key, []byte("/")...)

	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
