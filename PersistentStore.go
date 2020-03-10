package gunittesting

import (
	"errors"
	"gunittesting/domain"
	"gunittesting/interfaces"
)

type PersistentStore struct {
	store      interfaces.KeyValueStore
	serialiser interfaces.Serialiser
}

func NewPersistentStore(store interfaces.KeyValueStore, serialiser interfaces.Serialiser) *PersistentStore {
	return &PersistentStore{store: store, serialiser: serialiser}
}

// Stores a Frobnicator
func (ps *PersistentStore) Store(frob *domain.Frobnicator) error {
	// Validate
	err := validate(*frob)
	if err != nil {
		return err
	}

	// Serialise
	serialFrob, err := ps.serialiser.Dehydrate(frob)
	if err != nil {
		return errors.New("unable to serialise frob")
	}

	// Store
	err = ps.store.Store(frob.GetId(), serialFrob)
	if err != nil {
		return errors.New("unable to store frob")
	}

	return nil
}

// Retrieves a Frobnicator from Storage
func (ps *PersistentStore) Fetch(id string) (*domain.Frobnicator, error) {
	var frob *domain.Frobnicator

	// Fetch
	serialFrob, err := ps.store.Fetch(id)
	if err != nil {
		return frob, errors.New("error fetching from key/value store")
	}

	// Deserialise
	frob, err = ps.serialiser.Hydrate(serialFrob)
	if err != nil {
		return frob, errors.New("error deserialising frob")
	}

	return frob, nil
}

func validate(frob domain.Frobnicator) error {
	if frob.GetId() == "" {
		return errors.New("can't store a frob without an id")
	}

	return nil
}
