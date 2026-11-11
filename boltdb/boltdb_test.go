package boltdb

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDBListBuckets(t *testing.T) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("failed to get user home directory", err)
	}

	db, err := OpenDB(filepath.Join(homePath, ".cache", "trivy", "db", "trivy.db"))
	if err != nil {
		t.Fatal("failed to open DB:", err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			t.Error("failed to close DB:", err)
		}
	}()

	buckets, err := db.ListBuckets()
	if err != nil {
		t.Fatal("failed to list buckets:", err)
	}

	for _, buckName := range buckets {
		if bytes.HasPrefix(buckName, []byte("maven::")) {
			fmt.Printf("bucket: %s\n", string(buckName))
		}
	}
}
