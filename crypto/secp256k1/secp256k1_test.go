package secp256k1

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenKey(t *testing.T) {
	key, err := GenerateKey(rand.Reader)
	t.Logf("key: %x", key)
	assert.Equal(t, err, nil)
}
