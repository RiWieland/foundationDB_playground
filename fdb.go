package main

import (
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

type kvStore struct {
	instance fdb.Database
}

func initFDB() fdb.Database {
	fdb.MustAPIVersion(620)
	db := fdb.MustOpenDefault()
	db.Options().SetTransactionTimeout(60000) // 60,000 ms = 1 minute
	db.Options().SetTransactionRetryLimit(100)
	return db
}

func (db kvStore) initDirectory() directory.DirectorySubspace {
	directory, err := directory.CreateOrOpen(db.instance, []string{"accounts"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return directory

}
