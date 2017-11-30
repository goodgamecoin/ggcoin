package mptree

import (
	"github.com/goodgamecoin/ggcoin/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode(t *testing.T) {
	store := &MemStore{}
	n := newNode()
	var h common.Hash
	h[0], h[1], h[2] = 4, 5, 6
	err := n.insert(store, []byte{1, 2, 3}, &h)
	t.Logf("node %x", n)
	assert.Equal(t, nil, err)

	var h2 common.Hash
	h2[0], h2[1], h2[2], h2[3] = 4, 5, 6, 7
	err = n.insert(store, []byte{1, 2, 3, 4, 5, 6, 7}, &h2)

	cn, _, err := n.getChild(store, 3)
	assert.Equal(t, nil, err)
	t.Logf("child node %x", cn)

	t.Logf("memstore %x", store)
}
