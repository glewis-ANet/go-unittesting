package interfaces

import "gunittesting/domain"

type Serialiser interface {
	// Dehydrate
	//
	// Serialises a Frobnicator.
	Dehydrate(frob *domain.Frobnicator) (string, error)

	// Hydrate
	//
	// Deserialises a Frobnicator.
	Hydrate(frob string) (*domain.Frobnicator, error)
}
