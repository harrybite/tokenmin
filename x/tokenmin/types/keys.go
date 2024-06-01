package types

const (
	// ModuleName defines the module name
	ModuleName = "tokenmin"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tokenmin"
)

var (
	ParamsKey = []byte("p_tokenmin")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
