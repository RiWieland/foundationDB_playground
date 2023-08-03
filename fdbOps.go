package main

import (
	"log"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// create subspace/directory as separate structs
type kvStore struct {
	instance fdb.Database
	// subspaces []directory.DirectorySubspace
}

func initFdb() fdb.Database {
	fdb.MustAPIVersion(620)
	db := fdb.MustOpenDefault()
	db.Options().SetTransactionTimeout(60000) // 60,000 ms = 1 minute
	db.Options().SetTransactionRetryLimit(100)
	return db
}

// initialize directory for kvStore
func (db kvStore) addDirectorySub(name string) directory.DirectorySubspace {
	directorySub, err := directory.CreateOrOpen(db.instance, []string{name}, nil)
	if err != nil {
		log.Fatal(err)
	}

	/*
		var subspaces []directory.DirectorySubspace
		if len(db.subspaces) != 0 {
			subspaces = append(subspaces, db.subspaces...)
			db.subspaces = append(subspaces, directorySub)
		}*/
	return directorySub
}

// function writes the rectangle into specific sub
func (db kvStore) writeRect(r rectCoord) (err error) {
	_, err = db.instance.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Set(rectSub.Pack(tuple.Tuple{r.x0, r.x1, r.y0, r.y1}), []byte{})
		return
	})
	return
}

// Data model for the Files in KV-store:
// - Key path, fileType,Time
// - Value rect
func (db kvStore) writeImgWithCoor(f imgMeta, r rectCoord, time time.Duration) (err error) {

	rectKey := rectSub.Pack(tuple.Tuple{r.x0, r.x1, r.y0, r.y1})
	imgKey := imgSub.Pack(tuple.Tuple{f.path, f.fileType, int(time)})

	_, err = db.instance.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Set(rectKey, []byte(imgKey))
		return
	})
	return

}
