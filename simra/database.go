package simra

import "github.com/pankona/gomo-simra/simra/simlog"

// Databaser represents interface of database
type Databaser interface {
	Open(dirpath string) error
	Close()
	Put(key string, value interface{})
	Get(key string) interface{}
}

// Database is a container of Databaser
type Database struct {
	db Databaser
}

// OpenDB opens database connection
func OpenDB(databaser Databaser, dirpath string) (*Database, error) {
	err := databaser.Open(dirpath)
	if err != nil {
		simlog.Errorf("failed to open database. err = %s", err)
		return nil, err
	}
	return &Database{databaser}, nil
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
