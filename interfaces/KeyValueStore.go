package interfaces

type KeyValueStore interface {

	// Store
	//
	// Store the serialised data under the given key.
	Store(key string, data string) error

	// Fetch
	//
	// Fetch the serialised data for the given key.
	Fetch(key string) (string, error)
}
