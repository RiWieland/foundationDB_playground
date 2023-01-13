package main

import (
	"log"

	"errors"
	"fmt"
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

//var JennyAccount subspace.Subspace

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
	//	classKey := JennyAccount.Pack(tuple.Tuple{person, amount})
	fmt.Println(SCKey)

	_, err = t.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		if tr.Get(SCKey).MustGet() != nil {
			return // already signed up
		}

		classes := tr.GetRange(TimAccount.Sub(person), fdb.RangeOptions{Mode: fdb.StreamingModeWantAll}).GetSliceOrPanic()
		if len(classes) == 5 {
			err = errors.New("too many classes")
			return
		}

		tr.Set(SCKey, []byte(strconv.FormatInt(100, 10)))

		return
	})

	ret, err := t.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		ret = tr.Get(fdb.Key(SCKey)).MustGet()
		return
	})
	if err != nil {
		log.Fatalf("Unable to read FDB database value (%v)", err)
	}

	v := ret.([]byte)
	fmt.Printf("hello, %s\n", string(v))
	fmt.Printf("hello, %s\n", string(v))

	return

}
