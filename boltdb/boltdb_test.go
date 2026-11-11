package boltdb

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func getDBPath() (string, error) {
	userCachePath, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user cache directory: %w", err)
	}

	return filepath.Join(userCachePath, "trivy", "db", "trivy.db"), nil
}

func TestDBListBuckets(t *testing.T) {
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

func TestDBGetBucket(t *testing.T) {
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

	buckKeys, err := buck.ListKeys()
	if err != nil {
		t.Fatal("failed to get bucket keys:", err)
	}

	buckKey := []byte("com.fasterxml.jackson.core:jackson-core")

	found := false
	for _, k := range buckKeys {
		if bytes.Equal(k, buckKey) {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("failed to find jackson-core key")
	}

	buckVal, err := buck.GetValue(buckKey)
	if err != nil {
		t.Fatal("failed to get bucket value:", buckVal)
	}

	fmt.Printf("buck value: %s\n", string(buckVal))
}
