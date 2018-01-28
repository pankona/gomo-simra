package database

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/boltdb/bolt"
)

// Boltdb takes a role for interfaces of data store.
// implements simra.Databaser.
type Boltdb struct {
	db *bolt.DB
}

// Open opens new DB connection.
// Open will create a DB file under dirpath if not exist.
func (database *Boltdb) Open(dirpath string) error {
	db, err := bolt.Open(filepath.Join(dirpath, "db"), 0600, nil)
	if err != nil {
		return fmt.Errorf("failed to open database. error is: %s", err)
	}
	database.db = db
	err = db.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists([]byte("my_bucket"))
		if e != nil {
			log.Fatal(e)
			return nil
		}
		return nil
	})
	return err
}

// Close closes database.
// it is necessary to call this function after using database functions.
func (database *Boltdb) Close() {
	err := database.db.Close()
	if err != nil {
		log.Println(err)
	}
	database.db = nil
}

// Put puts a data to database.
// input must have ability to be casted into byte array.
func (database *Boltdb) Put(key string, value interface{}) {
	db := database.db
	if db == nil {
		log.Fatal("database is not opened yet.")
		return
	}

	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("my_bucket"))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		if str, ok := value.(string); ok {
			err = bucket.Put([]byte(key), []byte(str))
			if err != nil {
				log.Fatal(err)
				return nil
			}
		} else {
			log.Fatal("couldn't convert input to bytearray")
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

// Get returns put data.
func (database *Boltdb) Get(key string) interface{} {
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
