package main

import (
	"log"

	"errors"
	"strconv"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	/*
	  "github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	  "github.com/apple/foundationdb/bindings/go/src/fdb/tuple"

	  "fmt"
	  "strconv"
	  "errors"
	  "sync"
	  "math/rand"
	*/)

var initialBalanceTim = 100
var initialBalanceJenny = 300

var TimAccount subspace.Subspace
var JennyAccount subspace.Subspace

func main() {
	fdb.MustAPIVersion(620)
	db := fdb.MustOpenDefault()
	db.Options().SetTransactionTimeout(60000) // 60,000 ms = 1 minute
	db.Options().SetTransactionRetryLimit(100)

	accountsDir, err := directory.CreateOrOpen(db, []string{"accounts"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	TimAccount = accountsDir.Sub("class")
	//JennyAccount = accountsDir.Sub("attends")

	if err != nil {
		log.Fatalf("Unable to set FDB database value (%v)", err)
	}

	// Data Model for Key: ("AccountBalance", person, balance) = ""

	loadAccount(db, "Tim", 100)
	//loadAccount(db, "Jenny", 100)

}

func loadAccount(t fdb.Transactor, person string, amount int) (err error) {
	SCKey := TimAccount.Pack(tuple.Tuple{person, amount})
	classKey := JennyAccount.Pack(tuple.Tuple{person, amount})

	_, err = t.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		if tr.Get(SCKey).MustGet() != nil {
			return // already signed up
		}

		seats, err := strconv.ParseInt(string(tr.Get(classKey).MustGet()), 10, 64)
		if err != nil {
			return
		}
		if seats == 0 {
			err = errors.New("no remaining seats")
			return
		}

		classes := tr.GetRange(TimAccount.Sub(person), fdb.RangeOptions{Mode: fdb.StreamingModeWantAll}).GetSliceOrPanic()
		if len(classes) == 5 {
			err = errors.New("too many classes")
			return
		}

		tr.Set(classKey, []byte(strconv.FormatInt(seats-1, 10)))
		tr.Set(SCKey, []byte{})

		return
	})
	return
}
