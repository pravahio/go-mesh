package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func CreateNewAccount() (*ecdsa.PrivateKey, error) {
	prvKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return prvKey, nil
}
