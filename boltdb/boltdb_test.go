package boltdb

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func getDBPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	return filepath.Join(homePath, ".cache", "trivy", "db", "trivy.db"), nil
}

func TestListBuckets(t *testing.T) {
	dbPath, err := getDBPath()
	if err != nil {
		t.Fatal(err)
	}

	db, err := OpenDB(dbPath)
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

	found := false
	for _, buckName := range buckets {
		if bytes.HasPrefix(buckName, []byte("maven::")) {
			found = true
			break
		}
	}

	if !found {
		t.Error("failed to find bucket")
	}
}

func TestGetBucket(t *testing.T) {
	dbPath, err := getDBPath()
	if err != nil {
		t.Fatal("failure:", err)
	}

	db, err := OpenDB(dbPath)
	if err != nil {
		t.Fatal("failed to open DB:", err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			t.Error("failed to close DB:", err)
		}
	}()

	buck, err := db.GetBucket([]byte("maven::GitHub Security Advisory Maven"))
	if err != nil {
		t.Fatal("failed to get bucket:", err)
	}

	keys, err := buck.ListKeys()
	if err != nil {
		t.Fatal("failed to get bucket keys:", err)
	}

	for _, k := range keys {
		fmt.Printf("bucket key: %s\n", string(k))
	}
}
