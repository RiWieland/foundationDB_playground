package main

import (
	"fmt"
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
func (db kvStore) writeRect(r rectCoord) (f fdb.Key, err error) {
	rectKey := rectSub.Pack(tuple.Tuple{r.idx, r.x0, r.x1, r.y0, r.y1})

	_, err = db.instance.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Set(rectKey, []byte{})
		return
	})
	return rectKey, err
}

// Data model for the Files in KV-store:
// - Key path, Time
// - Value rect
func (db kvStore) writeImgWithCoor(f imgMeta, time time.Duration, r rectCoord) (Key fdb.Key, err error) {

	rectKey := rectSub.Pack(tuple.Tuple{r.x0, r.x1, r.y0, r.y1})
	imgKey := imgSub.Pack(tuple.Tuple{f.path, int(time)})
	recTest := []int{r.x0, r.x1, r.y0, r.y1}

	_, err = db.instance.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Set(imgKey, []byte(recTest))
		return
	})
	return imgKey, err

}

// clears Subspace
func (db kvStore) clearSub(FdbDir directory.DirectorySubspace) {

	_, _ = db.instance.Transact(func(tr fdb.Transaction) (interface{}, error) {
		tr.ClearRange(FdbDir)
		return nil, nil
	})
}

// Query of rectSub
func (db kvStore) queryRectSub() (ac []rectCoord, err error) {
	r, err := db.instance.ReadTransact(func(rtr fdb.ReadTransaction) (interface{}, error) {
		var rects []rectCoord
		ri := rtr.GetRange(rectSub, fdb.RangeOptions{}).Iterator()
		for ri.Advance() {

			kv := ri.MustGet()
			t, err := rectSub.Unpack(kv.Key)
			if err != nil {
				return nil, err
			}

			rectTemp := rectCoord{
				int(t[0].(int64)),
				int(t[1].(int64)),
				int(t[2].(int64)),
				int(t[3].(int64)),
				int(t[4].(int64)),
			}
			rects = append(rects, rectTemp)
		}
		return rects, nil
	})
	if err == nil {
		ac = r.([]rectCoord)
		fmt.Println(ac)
	}
	return
}

// Query of rectSub
func (db kvStore) queryImgSub() (ac []string, err error) {
	var classes []string
	r, err := db.instance.ReadTransact(func(rtr fdb.ReadTransaction) (interface{}, error) {

		ri := rtr.GetRange(imgSub, fdb.RangeOptions{}).Iterator()
		for ri.Advance() {

			kv := ri.MustGet()
			t, err := imgSub.Unpack(kv.Key)
			if err != nil {
				return nil, err
			}
			fmt.Println("t: key: ", t)

			v := kv.Value
			d := string(v[:])
			fmt.Println("d: ", d)

			fmt.Println("v: ", string(v))

			classes = append(classes, t[0].(string))

		}
		return classes, nil
	})
	if err == nil {
		ac = r.([]string)
		fmt.Println(ac)
	}
	return
}
