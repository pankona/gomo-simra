package simra

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/pankona/gomo-simra/simra/database"
)

type mock struct {
	m map[string]interface{}
}

func (db *mock) Open(dirpath string) error {
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
	db := OpenDB(&mock{}, filepath.Join("."))
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
	tmpdir, err := ioutil.TempDir(filepath.Join("."), "simradb_")
	if err != nil {
		t.Fatalf("failed to create temporary working directory. err: %s", err.Error())
	}
	defer func() {
		_ = os.RemoveAll(tmpdir)
	}()

	db := OpenDB(&database.Boltdb{}, filepath.Join(tmpdir))
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
