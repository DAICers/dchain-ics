package types

const (
	// ModuleName defines the module name
	ModuleName = "whiteboard"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_whiteboard"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	WhiteboardPixelKey      = "WhiteboardPixel/value/"
	WhiteboardPixelCountKey = "WhiteboardPixel/count/"
)

const (
	WhiteboardKey      = "Whiteboard/value/"
	WhiteboardCountKey = "Whiteboard/count/"
)
