package crypto

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowHash(t *testing.T) {
	data := []byte("This is a test")
	result := PowHash(data)
	r, _ := hex.DecodeString("a084f01d1437a09c6985401b60d43554ae105802c5f5d8a9b3253649c0be6605")
	assert.Equal(t, r, result[:])
}

func TestShaHash(t *testing.T) {
	data := []byte("")
	result := ShaHash(data)
	r, _ := hex.DecodeString("a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a")
	assert.Equal(t, r, result[:])
}

func TestGenerateKey(t *testing.T) {
	_, err := GenerateKey()
	//t.Logf("key: %x", key)
	assert.Equal(t, err, nil)
}
