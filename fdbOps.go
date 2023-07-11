package main

import (
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

type kvStore struct {
	instance fdb.Database
}

func (db kvStore) initFdb() kvStore {
	fdb.MustAPIVersion(620)
	db.instance = fdb.MustOpenDefault()
	db.instance.Options().SetTransactionTimeout(60000) // 60,000 ms = 1 minute
	db.instance.Options().SetTransactionRetryLimit(100)
	return db
}

func (db kvStore) initDirectory(name string) directory.DirectorySubspace {
	directory, err := directory.CreateOrOpen(db.instance, []string{name}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return directory

}
