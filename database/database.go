package database

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/xujiajun/nutsdb"
)

const (
	// BlocksBucket is the bucket to store blocks
	BlocksBucket = "blocksBucket"
	// TipKey is the key used to store the hash of the last block in the chain (its tip)
	TipKey = "l"
)

// Database embeds the original DB struct from NutsDB to utilize their methods and
// add custom ones on the same struct
type Database struct {
	*nutsdb.DB
}

// EmptyBucket checks if a given bucket is empty (no entries)
func (db *Database) EmptyBucket(bucket string) (bool, error) {
	var isEmpty bool

	err := db.View(
		func(tx *nutsdb.Tx) error {
			// Returns an err if bucket is empty
			_, err := tx.GetAll(bucket)
			if err != nil {
				isEmpty = true
				return err
			}
			return nil
		})

	return isEmpty, err
}

// Read from one key in specified bucket
func (db *Database) Read(bucket string, key []byte) ([]byte, error) {
	var value []byte
	err := db.View(
		func(tx *nutsdb.Tx) error {
			result, err := tx.Get(bucket, key)
			if err != nil {
				return err
			}
			value = result.Value
			return nil
		})
	return value, err
}

// Write one key/value pair to a specified bucket
func (db *Database) Write(bucket string, key, value []byte) error {
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
