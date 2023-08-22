package types

const (
	// ModuleName defines the module name.
	ModuleName = "tax"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for slashing.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_tax"
)

var (
	// TaxKey is the key to use for the keeper store.
	TaxKey    = []byte{0x00}
	ParamsKey = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
