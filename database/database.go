package database

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/xujiajun/nutsdb"
)

// For now, we use a single bucket storing all (unique) keys and values
const bucket = "bucket001"

func createDatabaseDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	databaseDirectory := cwd + "/db"

	_, err = ioutil.ReadDir(databaseDirectory)
	if err != nil {
		err := os.Mkdir(databaseDirectory, 0755)
		if err != nil {
			return "", err
		}
	}

	return databaseDirectory, nil
}

// Open is a wrapper around nutsDB.Open(), which opens and returns a new
// database instance.
func Open() *nutsdb.DB {
	dir, err := createDatabaseDirectory()
	if err != nil {
		log.Fatal(err)
	}
	opt := nutsdb.DefaultOptions
	opt.Dir = dir
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Write is a wrapper around nutsDB.Update(), which takes a key, and value and
// then writes to disk.
func Write(db *nutsdb.DB, key, value []byte) error {
	return db.Update(
		func(tx *nutsdb.Tx) error {
			err := tx.Put(bucket, key, value, 0)
			if err != nil {
				return err
			}
			return nil
		})
}

// Read in a wrapper around nutsDB.View(), which takes a key and returns its value.
func Read(db *nutsdb.DB, key []byte) (string, error) {
	var value string
	err := db.View(
		func(tx *nutsdb.Tx) error {
			result, err := tx.Get(bucket, key)
			if err != nil {
				return err
			}
			value = string(result.Value)
			return nil
		})
	return value, err
}
