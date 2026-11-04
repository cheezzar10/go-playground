package util

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestShaSum256(t *testing.T) {
	userHome, _ := os.UserHomeDir()
	pubKeyPath := filepath.Join(userHome, ".ssh", "id_rsa.pub")
	shasum, err := ShaSum256(pubKeyPath)
	if err != nil {
		t.Error("SHA256 calculation failed:", err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(shasum))
}
