package mptree

import (
	"github.com/goodgamecoin/ggcoin/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode1(t *testing.T) {
	store := &MemStore{}
	n := newNode()
	var h common.Hash
	h[0], h[1], h[2] = 4, 5, 6
	err := n.insert(store, []byte{1, 2, 3}, &h)
	assert.Equal(t, nil, err)
	n.save(store)

	var h2 common.Hash
	h2[0], h2[1], h2[2], h2[3] = 4, 5, 6, 7
	err = n.insert(store, []byte{1, 2, 4}, &h2)
	err = n.insert(store, []byte{1, 2, 5}, &h2)
	err = n.insert(store, []byte{1, 2, 6}, &h2)
	assert.Equal(t, nil, err)

	err = n.insert(store, []byte{1, 2}, &h2)
	assert.Equal(t, nil, err)

	n2, _, err := n.getChild(store, 2)

	t.Logf("node %s", n)
	t.Logf("node2 %s", n2)

	t.Logf("memstore %s", store)
}

func TestNode2(t *testing.T) {
	store := &MemStore{}
	n := newNode()
	var h common.Hash
	h[0], h[1], h[2] = 4, 5, 6
	err := n.insert(store, []byte{1, 2, 3}, &h)
	assert.Equal(t, nil, err)
	n.save(store)

	var h2 common.Hash
	h2[0], h2[1], h2[2], h2[3] = 4, 5, 6, 7
	err = n.insert(store, []byte{1, 2, 3, 4, 5, 6, 7}, &h2)
	assert.Equal(t, nil, err)

	err = n.insert(store, []byte{1, 2, 3, 4, 5, 6, 7, 8}, &h2)
	assert.Equal(t, nil, err)

	n2, _, err := n.getChild(store, 3)
	n3, _, err := n2.getChild(store, 7)

	t.Logf("node %s", n)
	t.Logf("node2 %s", n2)
	t.Logf("node3 %s", n3)

	t.Logf("memstore %s", store)

	err = n.insert(store, []byte{1, 2, 4}, &h2)
	assert.Equal(t, nil, err)
}
