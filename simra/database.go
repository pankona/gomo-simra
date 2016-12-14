package simra

import "fmt"

// Databaser represents interface of database
type Databaser interface {
	Open() error
	Close()
	Put(key string, value interface{})
	Get(key string) interface{}
}

// Database is a container of Databaser
type Database struct {
	db Databaser
}

// OpenDB opens database connection
func OpenDB(databaser Databaser) *Database {
	err := databaser.Open()
	if err != nil {
		_ = fmt.Errorf("failed to open database. err = %s", err)
		return nil
	}
	return &Database{databaser}
}

// Close closes database connection
func (database *Database) Close() {
	database.db.Close()
}

// Put stores a specified data to database
func (database *Database) Put(key string, value interface{}) {
	database.db.Put(key, value)
}

// Get fetches a data from database by specified key
func (database *Database) Get(key string) interface{} {
	return database.db.Get(key)
}
