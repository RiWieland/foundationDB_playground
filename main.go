package main

import (
	"log"

	"fmt"

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
	JennyAccount = accountsDir.Sub("attends")

	if err != nil {
		log.Fatalf("Unable to set FDB database value (%v)", err)
	}

	// Data Model for Key: ("AccountBalance", person, balance) = ""

	loadAccount(db, "Tim", 200)
	fetchAccount(db, "Tim", 200)

	//fetchAccount(db, "Jenny", 200)
	test, _ := listAllAccounts(db)
	fmt.Println(test)
}

func loadAccount(t fdb.Transactor, person string, amount int) (err error) {
	SCKey := TimAccount.Pack(tuple.Tuple{person, amount})
	fmt.Println(SCKey)
	// print encoding keys, more info.: https://forums.foundationdb.org/t/application-design-using-subspace-and-tuple/452

	_, err = t.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {

		tr.Set(SCKey, []byte{}) // we set an encoded key out of the Tuple of person and amount

		return
	})
	return
}

func fetchAccount(t fdb.Transactor, person string, amount int) (err error) {
	key := TimAccount.Pack(tuple.Tuple{person, amount})
	fmt.Println(key)

	ret, err := t.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		ret = tr.Get(fdb.Key(key)).MustGet()
		return
	})
	if err != nil {
		log.Fatalf("Unable to read FDB database value (%v)", err)
	}

	v := ret.([]byte)
	fmt.Printf("Amount: %s\n", string(v))

	return

}

func listAllAccounts(t fdb.Transactor) (ac []string, err error) {
	r, err := t.ReadTransact(func(rtr fdb.ReadTransaction) (interface{}, error) {
		var classes []string
		ri := rtr.GetRange(TimAccount, fdb.RangeOptions{}).Iterator()
		for ri.Advance() {
			kv := ri.MustGet()
			t, err := TimAccount.Unpack(kv.Key)
			if err != nil {
				return nil, err
			}
			classes = append(classes, t[0].(string))
		}
		return classes, nil
	})
	if err == nil {
		ac = r.([]string)
	}
	return
}

func dropAccount(t fdb.Transactor, person, amount int) (err error) {
	SCKey := TimAccount.Pack(tuple.Tuple{person, amount})

	_, err = t.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Clear(SCKey)
		return
	})
	return
}
