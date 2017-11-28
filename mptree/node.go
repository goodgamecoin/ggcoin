package mptree

import (
	"encoding/hex"
	//"github.com/rs/zerolog/log"
)

// https://github.com/ethereum/wiki/wiki/Patricia-Tree

// The node of a Merkle Patricia Tree
type node struct {
	path     []byte
	children map[byte][]byte
	data     map[byte][]byte
}

func (n *node) pathString() string {
	return hex.EncodeToString(n.path)
}

func (n *node) insert(kv KVStore, partialPath []byte, data []byte) error {
	if len(partialPath) < 2 {
		return ErrBadPathInsert
	}

	l := len(partialPath)
	path, tail := partialPath[:l-1], partialPath[l-1]
	if n == nil {
		*n = node{
			path: path,
			data: map[byte][]byte{tail: data},
		}
		return nil
	}

	pi := samePrefix(n.path, path)
	if pi < 1 {
		// They at least have one byte in common
		return ErrBadPathInsert
	} else if pi < len(n.path) {
		// We need to split current node path

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
