package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "parsa'sbestpicture"
	expectedPathKey := CASPathTransformFunc(key)
	expectedFilename := "2c9554a7b09dc3a6f19e1d3d0e3e089f44c86b06"
	expectedPathname := "2c955/4a7b0/9dc3a/6f19e/1d3d0/e3e08/9f44c/86b06"
	assert.Equal(t, expectedFilename, expectedPathKey.Filename)
	assert.Equal(t, expectedPathname, expectedPathKey.Pathname)
}

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg content")

		assert.NoError(t, s.writeStream(key, bytes.NewReader(data)))

		assert.Equal(t, s.Has(key), true)

		r, err := s.Read(key)
		assert.NoError(t, err)

		b, err := io.ReadAll(r)
		assert.NoError(t, err)

		assert.Equal(t, data, b)
	}
}

func TestStore_Delete(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	data := []byte("some jpg content")
	key := "myspecialpicture"

	assert.NoError(t, s.writeStream(key, bytes.NewReader(data)))
	assert.NoError(t, s.Delete(key))
	assert.Equal(t, s.Has(key), false)
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	return NewStore(opts)
}

func teardown(t *testing.T, store *Store) {
	assert.NoError(t, store.clear())
}
