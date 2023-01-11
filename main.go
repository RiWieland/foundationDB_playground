package main

import (
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
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

func main() {
	fdb.MustAPIVersion(620)
	db := fdb.MustOpenDefault()
	db.Options().SetTransactionTimeout(60000) // 60,000 ms = 1 minute
	db.Options().SetTransactionRetryLimit(100)

	accountsDir, err := directory.CreateOrOpen(db, []string{"accounts"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	TimAccount := accountsDir.Sub("Tim")
	JennyAccount = accountsDir.Sub("Jenny")

	if err != nil {
		log.Fatalf("Unable to set FDB database value (%v)", err)
	}

	// Data Model for Key: ("AccountBalance", person, balance) = ""

	loadAccount(fdb.Transaction{}, TimAccount, "Tim", 100)

}

func loadAccount(t fdb.Transactor, Account directory.DirectorySubspace, person string, amount int) (err error) {
	_, err = t.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Set(Account.Pack(tuple.Tuple{person, amount}), []byte{})
		return
	})
	return
}
