package util

import (
	"crypto/sha256"
	"io"
	"os"
)

func ShaSum256(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	buf := make([]byte, 4096)
	hash := sha256.New()
	for {
		n, err := f.Read(buf)
		hash.Write(buf[:n])
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
	}

	return hash.Sum(nil), nil
}
