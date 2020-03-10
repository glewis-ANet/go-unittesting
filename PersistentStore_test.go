package gunittesting

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gunittesting/domain"
	"gunittesting/mocks"
	"testing"
)

// Happy path
// Mock verification
func TestStore_ValidFrobnicator_Success(t *testing.T) {
	// Arrange
	frob := &domain.Frobnicator{Id: "key", Bar: "value"}
	mockSerialiser := &mocks.Serialiser{}
	mockStore := &mocks.KeyValueStore{}
	store := &PersistentStore{serialiser: mockSerialiser, store: mockStore}

	mockSerialiser.On("Dehydrate", frob).Return("weetbix", nil).Once()
	mockStore.On("Store", "key", "weetbix").Return(nil).Once()

	// Act
	err := store.Store(frob)

	// Assert
	mockStore.AssertExpectations(t)
	assert.True(t, err == nil)
}

// Simple failure test
func TestStore_InvalidFrobnicator_Error(t *testing.T) {
	// Arrange
	frob := &domain.Frobnicator{Id: "", Bar: "value"}
	store := &PersistentStore{serialiser: nil, store: nil}

	// Act
	err := store.Store(frob)

	// Assert
	assert.True(t, err != nil)
}

// Serialisation error path
func TestStore_SerialisationFailure_Error(t *testing.T) {
	// Arrange
	frob := &domain.Frobnicator{Id: "key", Bar: "value"}
	mockSerialiser := &mocks.Serialiser{}
	store := &PersistentStore{serialiser: mockSerialiser, store: nil}

	mockSerialiser.On("Dehydrate", frob).Return("", errors.New("serialiser error")).Once()

	// Act
	err := store.Store(frob)

	// Assert
	mockSerialiser.AssertExpectations(t)
	assert.True(t, err != nil)
}

// Storage error path
func TestStore_StorageFailure_Error(t *testing.T) {
	// Arrange
	frob := &domain.Frobnicator{Id: "key", Bar: "value"}
	mockSerialiser := &mocks.Serialiser{}
	mockStore := &mocks.KeyValueStore{}
	store := &PersistentStore{serialiser: mockSerialiser, store: mockStore}

	mockSerialiser.On("Dehydrate", frob).Return("weetbix", nil).Once()
	mockStore.On("Store", "key", "weetbix").Return(errors.New("storage error")).Once()

	// Act
	err := store.Store(frob)

	// Assert
	mockStore.AssertExpectations(t)
	assert.True(t, err != nil)
}

// Happy path
func TestFetch_ValidId_Success(t *testing.T) {
	// Arrange
	key := "key"
	serialised := "weetbix"
	frob := &domain.Frobnicator{Id: key, Bar: "value"}
	mockSerialiser := &mocks.Serialiser{}
	mockStore := &mocks.KeyValueStore{}
	store := &PersistentStore{serialiser: mockSerialiser, store: mockStore}

	mockStore.On("Fetch", key).Return(serialised, nil).Once()
	mockSerialiser.On("Hydrate", serialised).Return(frob, nil).Once()

	// Act
	result, err := store.Fetch(key)

	// Assert
	mockStore.AssertExpectations(t)
	assert.True(t, result.GetId() == key)
	assert.True(t, err == nil)
}
