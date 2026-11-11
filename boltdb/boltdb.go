// Package boltdb provides a thin wrapper around go.etcd.io/bbolt key/value store.
package boltdb

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

// DB wraps a Bolt database.
type DB struct {
	db *bolt.DB
}

// OpenDB opens a Bolt database at path.
func OpenDB(path string) (*DB, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

// ListBuckets returns copies of all top-level bucket names.
func (db *DB) ListBuckets() ([][]byte, error) {
	var buckets [][]byte

	err := db.db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(bucketName []byte, _ *bolt.Bucket) error {
			bucketNameCopy := make([]byte, len(bucketName), cap(bucketName))
			copy(bucketNameCopy, bucketName)
			buckets = append(buckets, bucketNameCopy)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return buckets, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func ReadBucket(db *bolt.DB, bucketId string) error {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketId))
		if b == nil {
			return fmt.Errorf("boltdb: bucket %q not found", string(bucketId))
		}
		return b.ForEach(func(k, v []byte) error {
			fmt.Printf("entry key: %s\n", string(k))
			return nil
		})
	})
	if err != nil {
		return fmt.Errorf("boltdb: read bucket %q: %w", bucketId, err)
	}

	return nil
}
