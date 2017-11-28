package mptree

import (
	"errors"
)

var (
	KVStoreNotFound = errors.New("Key not found in KV store")

	ErrBadPathInsert = errors.New("Bad path when inserting")

	MemStoreIvalidType = errors.New("Got invalid type in KV store")
)
