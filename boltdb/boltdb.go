// Package boltdb provides a thin wrapper around go.etcd.io/bbolt key/value store.
package boltdb

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func Open(path string) (*bolt.DB, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
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

func ReadBuckets(db *bolt.DB) error {
	return db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, bucket *bolt.Bucket) error {
			fmt.Printf("bucket name: %s\n", string(name))
			return nil
		})
	})
}
