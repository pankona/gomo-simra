package simra

import (
	"log"

	"github.com/boltdb/bolt"
)

// Database takes a role for interfaces of data store.
type Database struct {
	db *bolt.DB
}

// Open opens database.
// it is necessary to call this function before using database functions.
func (database *Database) Open() {
	db, err := bolt.Open("db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	database.db = db
}

// Close closes database.
// it is necessary to call this function after using database functions.
func (database *Database) Close() {
	database.db.Close()
	database.db = nil
}

// Put puts a data to database.
// input must have ability to be cated into byte array.
func (database *Database) Put(key string, value interface{}) {
	db := database.db
	if db == nil {
		log.Fatal("database is not opened yet.")
		return
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("my_bucket"))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		if byteArray, ok := value.([]byte); ok {
			err = bucket.Put([]byte(key), byteArray)
			if err != nil {
				log.Fatal(err)
				return nil
			}
		} else {
			log.Fatal("couldn't convert input to bytearray")
		}
		return nil
	})
}

// Get returns put data.
func (database *Database) Get(key string) interface{} {
	db := database.db
	if db == nil {
		log.Fatal("database is not opened yet.")
		return nil
	}
	var value interface{}
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("my_bucket"))
		if bucket == nil {
			log.Fatal("bucket not found")
			return nil
		}
		value = bucket.Get([]byte(key))
		return nil
	})
	if err != nil {
		return nil
	}
	return value
}
