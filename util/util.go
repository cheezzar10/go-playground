package util

import (
	"crypto/sha256"
)

func ShaSum256(data []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil), nil
}
