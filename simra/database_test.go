package simra

import (
	"log"
	"testing"
)

type mock struct {
	m map[string]interface{}
}

func (db *mock) Open() error {
	if db.m == nil {
		db.m = map[string]interface{}{}
	}
	return nil
}

func (db *mock) Close() {
	db.m = nil
}

func (db *mock) Put(key string, value interface{}) {
	if db.m == nil {
		log.Fatal("database is not opened yet")
		return
	}
	db.m[key] = value
}

func (db *mock) Get(key string) interface{} {
	if db.m == nil {
		log.Fatal("database is not opened yet")
		return nil
	}
	return db.m[key]
}

func TestMock(t *testing.T) {
	db := OpenDB(&mock{})
	if db == nil {
		t.Error("failed to open database")
	}
	defer db.Close()

	db.Put("key1", "value1")
	value1 := db.Get("key1")
	if value1 != "value1" {
		t.Error("failed to open database")
	}
}

func TestBolt(t *testing.T) {
	db := OpenDB(&boltdb{})
	if db == nil {
		t.Error("failed to open database")
	}
	defer db.Close()

	db.Put("key1", "value1")
	fetched := db.Get("key1")
	if bytes, ok := fetched.([]uint8); ok {
		v := string(bytes)
		if v != "value1" {
			t.Errorf("unexpected fetched value. expected: %s, fetched: %s", "value1", v)
		}
	}
}
