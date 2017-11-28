package mptree

import (
	"encoding/hex"
	"sync"
)

type KVStore interface {
	KVReader
	KVWriter
}

type KVReader interface {
	Get(key []byte) (value []byte, err error)
	Has(key []byte) (bool, error)
}

type KVWriter interface {
	Put(key, value []byte) error
	Del(key []byte) error
}

type MemStore struct {
	sync.Map
}

func (m *MemStore) Get(key []byte) (value []byte, err error) {
	val, ok := m.Load(bytesToStr(key))
	if !ok {
		return nil, KVStoreNotFound
	}
	if b, ok := val.([]byte); !ok {
		return nil, MemStoreIvalidType
	} else {
		return b, nil
	}
}

func (m *MemStore) Has(key []byte) (bool, error) {
	_, ok := m.Load(bytesToStr(key))
	return ok, nil
}

func (m *MemStore) Put(key, value []byte) error {
	m.Store(bytesToStr(key), value)
	return nil
}

func (m *MemStore) Del(key []byte) error {
	m.Delete(bytesToStr(key))
	return nil
}

func bytesToStr(b []byte) string {
	return hex.EncodeToString(b)
}
