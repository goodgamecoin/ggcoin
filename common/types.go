package common

import (
	"encoding/hex"
	"math/big"
)

const (
	HashLength = 32
)

type Hash [HashLength]byte

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func BytesToHashPtr(b []byte) *Hash {
	var h Hash
	h.SetBytes(b)
	return &h
}

func StringToHash(s string) Hash { return BytesToHash([]byte(s)) }

//func BigToHash(b *big.Int) Hash  { return BytesToHash(b.Bytes()) }
//func HexToHash(s string) Hash    { return BytesToHash(FromHex(s)) }

func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}
	copy(h[HashLength-len(b):], b)
}

func (h *Hash) Str() string { return string(h[:]) }

func (h *Hash) String() string { return hex.EncodeToString(h[:]) }

func (h *Hash) Bytes() []byte { return h[:] }

func (h *Hash) Big() *big.Int { return new(big.Int).SetBytes(h[:]) }
