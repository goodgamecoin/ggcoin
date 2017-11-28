package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/goodgamecoin/ggcoin/crypto/cryptonight"
	"github.com/goodgamecoin/ggcoin/crypto/secp256k1"
	"golang.org/x/crypto/sha3"
)

// The generic cryptographically secure hash function
func ShaHash(d []byte) [32]byte {
	return sha3.Sum256(d)
}

// The PoW hash used for mining
func PowHash(d []byte) [32]byte {
	return cryptonight.Hash(d)
}

// Generate secp256k1 key with OS provided random function
func GenerateKey() (*ecdsa.PrivateKey, error) {
	return secp256k1.GenerateKey(rand.Reader)
}
