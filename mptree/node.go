package mptree

import (
	"bytes"
	"encoding/hex"
	"github.com/goodgamecoin/ggcoin/common"
	"github.com/goodgamecoin/ggcoin/common/rlp"
	"github.com/goodgamecoin/ggcoin/crypto"
	"strconv"
	//"github.com/rs/zerolog/log"
)

// https://github.com/ethereum/wiki/wiki/Patricia-Tree

// The node of a Merkle Patricia Tree
type node struct {
	path     []byte
	data     map[byte]*common.Hash
	children map[byte]*common.Hash
}

func newNode() *node {
	return &node{
		data:     make(map[byte]*common.Hash),
		children: make(map[byte]*common.Hash),
	}
}

func (n *node) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{[PATH:")
	buffer.WriteString(hex.EncodeToString(n.path))
	buffer.WriteString("];[DATA: ")
	for i := 0; i < 255; i++ {
		if h := n.data[byte(i)]; h != nil {
			buffer.WriteString(strconv.Itoa(i))
			buffer.WriteString("->")
			buffer.WriteString(h.String())
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("];[CHILDREN: ")
	for i := 0; i < 255; i++ {
		if h := n.children[byte(i)]; h != nil {
			buffer.WriteString(strconv.Itoa(i))
			buffer.WriteString("->")
			buffer.WriteString(h.String())
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("]}")
	return buffer.String()
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
func (n *node) insert(kv KVStore, partialPath []byte, insertHash *common.Hash) error {
	if len(partialPath) < 2 {
		return ErrBadPathInsert
	}

	l := len(partialPath)
	insertPath, tail := partialPath[:l-1], partialPath[l-1]
	if len(n.path) == 0 {
		n.path, n.data[tail] = insertPath, insertHash
		return nil
	}

	pi := samePrefix(n.path, insertPath)
	if pi < 1 {
		// They must have at least one byte in common
		return ErrBadPathInsert
	} else if pi < len(n.path) {
		// We need to split current node path
		child1 := &node{
			path:     n.path[pi:],
			data:     n.data,
			children: n.children,
		}
		h, err := child1.save(kv)
		if err != nil {
			return err
		}
		n.path = n.path[:pi]
		n.data = make(map[byte]*common.Hash)
		n.children = make(map[byte]*common.Hash)
		n.children[child1.path[0]] = h

		if len(insertPath) == len(n.path) { // Store data at current node
			n.data[tail] = insertHash
		} else { // Need to create another child node
			child2 := newNode()
			child2.path = insertPath[pi:]
			child2.data[tail] = insertHash
			h, err := child1.save(kv)
			if err != nil {
				return err
			}
			n.children[child2.path[0]] = h
		}
	} else {
		if len(insertPath) == len(n.path) {
			n.data[tail] = insertHash
			return nil // We're done here in this case
		}
		// In this case, we create a new node as a child node
		index := insertPath[pi]
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
		if err := childNode.insert(kv, partialPath[pi:], insertHash); err != nil {
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

func (n *node) remove(kv KVStore, path []byte) error {
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
