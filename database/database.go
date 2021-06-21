package database

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/xujiajun/nutsdb"
)

// For now, we use a single bucket storing all (unique) keys and values
const bucket = "bucket001"

// Database embeds the original DB struct from NutsDB to utilize their methods and
// add custom ones on the same struct
type Database struct {
	*nutsdb.DB
}

func (db *Database) Read(key []byte) (string, error) {
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

func (db *Database) Write(key, value []byte) error {
	return db.Update(
		func(tx *nutsdb.Tx) error {
			err := tx.Put(bucket, key, value, 0)
			if err != nil {
				return err
			}
			return nil
		})
}

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
// custom database instance, embedding the nutsDB database struct.
func Open() *Database {
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
	return &Database{db}
}
