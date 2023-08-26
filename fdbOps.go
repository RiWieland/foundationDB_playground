package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
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
func (db kvStore) writeImgWithCoor(f img, time time.Duration, r rectCoord) (Key fdb.Key, err error) {

	imgKey := imgSub.Pack(tuple.Tuple{f.path, int(time)})

	_, err = db.instance.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Set(imgKey, []byte(strconv.FormatInt(r.idx, 10)))
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

// Query of rectSub to return all rects in the sub in a slice
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
				t[0].(int64),
				t[1].(int64),
				t[2].(int64),
				t[3].(int64),
				t[4].(int64),
			}
			rects = append(rects, rectTemp)
		}
		return rects, nil
	})
	if err == nil {
		ac = r.([]rectCoord)
	}
	return
}

// Query of ImgSub
// The query returns the key and value for the whole imageSub
func (db kvStore) queryImgSub() (ac []img, err error) {
	var imgScan []img
	var imgReturn img
	//var rectReturn rect
	r, err := db.instance.ReadTransact(func(rtr fdb.ReadTransaction) (interface{}, error) {

		ri := rtr.GetRange(imgSub, fdb.RangeOptions{}).Iterator()
		for ri.Advance() {

			kv := ri.MustGet()
			t, err := imgSub.Unpack(kv.Key)
			if err != nil {
				return nil, err
			}

			v := string(kv.Value)
			ret := convertToInt(v)
			s, _ := db.queryRect(int64(ret))
			imgReturn := imgReturn.initImgObj(t[0].(string))
			imgReturn.rect = s
			imgScan = append(imgScan, imgReturn)

		}
		return imgScan, nil

	})
	if err == nil {
		ac = r.([]img)
		fmt.Println(ac)
	}
	return

}

// query for individual Rect by idx
// ATM: queries the subspace and filters on the specific
func (db kvStore) queryRect(idx int64) (rectCoord, error) {
	r, _ := db.queryRectSub()
	var rectReturn rectCoord
	for _, rect := range r {
		if idx == int64(rect.idx) {
			rectReturn = rect
		}
	}
	if rectReturn != (rectCoord{}) {
		return rectReturn, nil

	} else {
		return rectReturn, errors.New("zero Value returned from query rectSub")
	}

}
