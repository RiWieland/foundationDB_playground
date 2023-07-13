package main

import (
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

// create subspace/directory as separate structs
type kvStore struct {
	instance  fdb.Database
	subspaces []directory.DirectorySubspace
}

func (db kvStore) initFdb() kvStore {
	fdb.MustAPIVersion(620)
	db.instance = fdb.MustOpenDefault()
	db.instance.Options().SetTransactionTimeout(60000) // 60,000 ms = 1 minute
	db.instance.Options().SetTransactionRetryLimit(100)
	return db
}

// initialize directory for kvStore
func (db kvStore) addDirectorySub(name string) {
	directorySub, err := directory.CreateOrOpen(db.instance, []string{name}, nil)
	if err != nil {
		log.Fatal(err)
	}

	var subspaces []directory.DirectorySubspace
	if len(db.subspaces) != 0 {
		subspaces = append(subspaces, db.subspaces...)
		db.subspaces = append(subspaces, directorySub)
	}
}
