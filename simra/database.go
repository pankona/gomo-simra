package simra

// Database takes a role for interfaces of data store.
type Database struct{}

func (db *Database) Put() {
}

func (db *Database) Get() interface{} {
	return nil
}
