package cryptonight

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	data := []byte("This is a test")
	result := Hash(data)
	r, _ := hex.DecodeString("a084f01d1437a09c6985401b60d43554ae105802c5f5d8a9b3253649c0be6605")
	assert.Equal(t, r, result[:])
}
