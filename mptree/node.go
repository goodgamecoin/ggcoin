package mptree

import (
	"bytes"
	"encoding/hex"
	"github.com/goodgamecoin/ggcoin/common"
	"github.com/goodgamecoin/ggcoin/common/rlp"
	"github.com/goodgamecoin/ggcoin/crypto"
	//"github.com/rs/zerolog/log"
)

// https://github.com/ethereum/wiki/wiki/Patricia-Tree

// The node of a Merkle Patricia Tree
type node struct {
	path     []byte
	children map[byte]*common.Hash
	data     map[byte]*common.Hash
}

func newNode() *node {
	return &node{
		children: make(map[byte]*common.Hash),
		data:     make(map[byte]*common.Hash),
	}
}

func (n *node) hasChild(index byte) bool {
	return n.children[index] != nil
}

func (n *node) setChild(index byte, hash *common.Hash) {
	n.children[index] = hash
}

func (n *node) getChild(kv KVStore, index byte) (*node, *common.Hash, error) {
	if childPtr := n.children[index]; childPtr == nil {
		return nil, nil, NodeChildNotFound
	} else {
		val, err := kv.Get(childPtr.Bytes())
		if err != nil {
			return nil, nil, err
		}
		var cnode node
		b := bytes.NewBuffer(val)
		if err := rlp.Decode(b, &cnode); err != nil {
			return nil, nil, err
		}
		return &cnode, childPtr, nil
	}
}

func (n *node) save(kv KVStore) (*common.Hash, error) {
	var b bytes.Buffer
	if err := rlp.Encode(&b, n); err != nil {
		return nil, err
	}
	d := b.Bytes()
	hash := crypto.ShaHash(d)
	return &hash, kv.Put(hash.Bytes(), d)
}

func (n *node) pathString() string {
	return hex.EncodeToString(n.path)
}

// This function changes the content of the node, it is the caller's responsibility
// to delete the old hash-value entry and add the new entry
func (n *node) insert(kv KVStore, partialPath []byte, hash *common.Hash) error {
	if len(partialPath) < 2 {
		return ErrBadPathInsert
	}

	l := len(partialPath)
	path, tail := partialPath[:l-1], partialPath[l-1]
	if len(n.path) == 0 {
		n.path, n.data[tail] = path, hash
		return nil
	}

	pi := samePrefix(n.path, path)
	if pi < 1 {
		// They must have at least one byte in common
		return ErrBadPathInsert
	} else if pi < len(n.path) {
		// We need to split current node path

	} else {
		if len(path) == len(n.path) {
			n.data[tail] = hash
			return nil // We're done here in this case
		}

		index := path[pi]
		childNode := newNode()
		var childHash *common.Hash
		if n.hasChild(index) {
			node, hash, err := n.getChild(kv, index)
			if err != nil {
				return err
			}
			childNode = node
			childHash = hash
		}
		if err := childNode.insert(kv, partialPath[pi:], hash); err != nil {
			return err
		}
		if h, err := childNode.save(kv); err != nil {
			return err
		} else {
			n.setChild(index, h)
		}
		if childHash != nil {
			return kv.Del(childHash.Bytes())
		}
	}
	return nil
}

func samePrefix(p1, p2 []byte) int {
	i := 0
	for ; i < len(p1) && i < len(p2); i++ {
		if p1[i] != p2[i] {
			return i
		}
	}
	return i
}
