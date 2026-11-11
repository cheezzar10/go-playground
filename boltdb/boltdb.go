// Package boltdb provides a thin wrapper around go.etcd.io/bbolt key/value store.
package boltdb

import (
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
		return tx.ForEach(func(buckName []byte, _ *bolt.Bucket) error {
			buckets = append(buckets, append([]byte{}, buckName...))
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

type Bucket struct {
	db   *DB
	name []byte
}

func (db *DB) GetBucket(buckName []byte) (*Bucket, error) {
	var bucket *Bucket
	err := db.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(buckName)
		if b == nil {
			return nil
		}

		bucket = &Bucket{db, buckName}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func (b *Bucket) ListKeys() ([][]byte, error) {
	keys := [][]byte{}

	err := b.db.db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket(b.name)
		if buck == nil {
			// handle
			return nil
		}

		return buck.ForEach(func(key, val []byte) error {
			keys = append(keys, append([]byte{}, key...))
			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return keys, nil
}

// GetValue returns a copy of the value associated with key.
func (b *Bucket) GetValue(key []byte) ([]byte, error) {
	var value []byte

	err := b.db.db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket(b.name)
		if buck == nil {
			return nil
		}

		val := buck.Get(key)
		if val != nil {
			value = append([]byte{}, val...)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}
