package mptree

import (
	"bytes"
	"github.com/goodgamecoin/ggcoin/common"
	"github.com/goodgamecoin/ggcoin/common/rlp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndecode(t *testing.T) {
	na := &node{
		path:     []byte{1, 2, 3},
		children: map[byte]*common.Hash{4: common.BytesToHashPtr([]byte{1, 2}), 5: common.BytesToHashPtr([]byte{3, 4})},
		data:     map[byte]*common.Hash{6: common.BytesToHashPtr([]byte{5, 6, 7}), 7: common.BytesToHashPtr([]byte{8, 9})},
	}
	var buffer bytes.Buffer
	assert.Equal(t, nil, rlp.Encode(&buffer, na))
	t.Logf("buffer: %x", buffer.Bytes())

	var n2 node
	err := rlp.Decode(&buffer, &n2)
	assert.Equal(t,
		nil, err,
		na.path, n2.path,
		na.children[4], n2.children[4],
		na.data[6], n2.data[6])

	nb := &node{
		path: []byte{1, 2, 3},
		data: map[byte]*common.Hash{16: common.BytesToHashPtr([]byte{255, 34, 7}), 7: common.BytesToHashPtr([]byte{8, 9})},
	}
	buffer.Reset()
	assert.Equal(t, nil, rlp.Encode(&buffer, nb))
	t.Logf("buffer: %x", buffer.Bytes())
	err = rlp.Decode(&buffer, &n2)
	assert.Equal(t,
		nil, err,
		na.path, n2.path,
		na.children[4], n2.children[4],
		na.data[6], n2.data[6])
}
