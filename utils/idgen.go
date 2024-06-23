package utils

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
const length = 6

func GenerateID() (string, error) {
	b := make([]byte, length)
	
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		b[i] = charset[n.Int64()]
	}

	return string(b), nil
}
