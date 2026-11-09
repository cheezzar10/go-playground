package boltdb

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("failed to get user home directory", err)
	}
	dbPath := filepath.Join(homePath, ".cache", "trivy", "fanal", "fanal.db")

	db, err := Open(dbPath)
	if err != nil {
		t.Fatal("failed to open db:", err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			t.Error("failed to close DB:", err)
		}
	}()

	err = ReadBuckets(db)
	if err != nil {
		t.Fatal("failed to read buckets:", err)
	}

	// 'blob', 'artifact' buckets
	err = ReadBucket(db, "blob")
	if err != nil {
		t.Fatal("failed to read bucket:", err)
	}
}
