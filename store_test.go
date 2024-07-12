package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "parsa'sbestpicture"
	expectedPathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "2c9554a7b09dc3a6f19e1d3d0e3e089f44c86b06"
	expectedPathname := "2c955/4a7b0/9dc3a/6f19e/1d3d0/e3e08/9f44c/86b06"
	assert.Equal(t, expectedOriginalKey, expectedPathKey.FullPath)
	assert.Equal(t, expectedPathname, expectedPathKey.Pathname)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	data := []byte("some jpg content")
	key := "myspecialpicture"

	assert.NoError(t, s.writeStream(key, bytes.NewReader(data)))

	r, err := s.Read(key)
	assert.NoError(t, err)

	b, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Equal(t, data, b)
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
}
