package types

const (
	// ModuleName defines the module name
	ModuleName = "treasury"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

    // QuerierRoute defines the module's query routing key
    QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_treasury"

    
)

var (
	// Keys for store prefixes
	GenesisStateKey = []byte{0x01} // a key for the genesis state
)


func KeyPrefix(p string) []byte {
    return []byte(p)
}
