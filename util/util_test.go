package util

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestShaSum256(t *testing.T) {
	userHome, _ := os.UserHomeDir()
	pubKeyPath := filepath.Join(userHome, ".ssh", "id_rsa.pub")
	pubKeyFile, err := os.Open(pubKeyPath)
	if err != nil {
		t.Error("failed to open file:", err)
	}

	defer pubKeyFile.Close()

	pubKeyReader := bufio.NewReader(pubKeyFile)

	_, err = pubKeyReader.ReadString(' ')
	if err != nil {
		t.Error("failed to read public key preamble:", err)
	}

	pubKeyBase64, err := pubKeyReader.ReadString(' ')
	if err != nil {
		t.Error("failed to read public key:", err)
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(pubKeyBase64))
	if err != nil {
		t.Error("public key decoding failed:", err)
	}

	shasum, err := ShaSum256(pubKeyBytes)
	if err != nil {
		t.Error("SHA256 calculation failed:", err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(shasum))
}
