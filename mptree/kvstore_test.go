package mptree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemStore(t *testing.T) {
	var store MemStore
	one := []byte{1}
	_, err := store.Get(one)
	assert.Equal(t, KVStoreNotFound, err)

	store.Put(one, []byte{2})
	re, err := store.Get(one)
	assert.Equal(t, []byte{2}, re)
}
