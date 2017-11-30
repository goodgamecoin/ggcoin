package mptree

import (
	"errors"
)

var (
	KVStoreNotFound = errors.New("Key not found in KV store")

	NodeChildNotFound = errors.New("Node child not found")

	ErrBadPathInsert = errors.New("Bad path when inserting")

	ErrBadEncodedBinary = errors.New("Bad rlp encoding")

	MemStoreIvalidType = errors.New("Got invalid type in KV store")
)
