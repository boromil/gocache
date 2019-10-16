package store

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBigcache(t *testing.T) {
	// Given
	client := &MockBigcacheClientInterface{}

	// When
	store := NewBigcache(client, nil)

	// Then
	assert.IsType(t, new(BigcacheStore), store)
	assert.Equal(t, client, store.client)
	assert.IsType(t, new(Options), store.options)
}

func TestBigcacheGet(t *testing.T) {
	// Given
	cacheKey := "my-key"
	cacheValue := []byte("my-cache-value")

	client := &MockBigcacheClientInterface{}
	client.On("Get", cacheKey).Return(cacheValue, nil)

	store := NewBigcache(client, nil)

	// When
	value, err := store.Get(cacheKey)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, cacheValue, value)
}

func TestBigcacheSet(t *testing.T) {
	// Given
	cacheKey := "my-key"
	cacheValue := []byte("my-cache-value")

	client := &MockBigcacheClientInterface{}
	client.On("Set", cacheKey, cacheValue).Return(nil)

	store := NewBigcache(client, nil)

	// When
	err := store.Set(cacheKey, cacheValue, nil)

	// Then
	assert.Nil(t, err)
}

func TestBigcacheDelete(t *testing.T) {
	// Given
	cacheKey := "my-key"

	client := &MockBigcacheClientInterface{}
	client.On("Delete", cacheKey).Return(nil)

	store := NewBigcache(client, nil)

	// When
	err := store.Delete(cacheKey)

	// Then
	assert.Nil(t, err)
}

func TestBigcacheDeleteWhenError(t *testing.T) {
	// Given
	expectedErr := errors.New("Unable to delete key")

	cacheKey := "my-key"

	client := &MockBigcacheClientInterface{}
	client.On("Delete", cacheKey).Return(expectedErr)

	store := NewBigcache(client, nil)

	// When
	err := store.Delete(cacheKey)

	// Then
	assert.Equal(t, expectedErr, err)
}

func TestBigcacheGetType(t *testing.T) {
	// Given
	client := &MockBigcacheClientInterface{}

	store := NewBigcache(client, nil)

	// When - Then
	assert.Equal(t, BigcacheType, store.GetType())
}
