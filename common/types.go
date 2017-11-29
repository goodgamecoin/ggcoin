package common

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
