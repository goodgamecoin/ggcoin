package secp256k1

import (
	"crypto/ecdsa"
	"io"
)

func GenerateKey(rand io.Reader) (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(secp256k1Params, rand)
}
