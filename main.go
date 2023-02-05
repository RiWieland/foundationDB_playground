package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

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

// To do:
// - transfer data between routines -> channnel /chan
// - make sure every person only got one account
// - implement restrictions:
// -  - no negative vales
// - overwriting keys in key-value store OR timestamp?

var initialBalanceTim = 100
var initialBalanceJenny = 300

var TimAccount subspace.Subspace
var JennyAccount subspace.Subspace

type personalAccount struct {
	name   string
	amount int64
}

type accountList struct {
	bank    string
	members []personalAccount
}

func main() {
	fdb.MustAPIVersion(620)
	db := fdb.MustOpenDefault()
	db.Options().SetTransactionTimeout(60000) // 60,000 ms = 1 minute
	db.Options().SetTransactionRetryLimit(100)

	accountsDir, err := directory.CreateOrOpen(db, []string{"accounts"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	TimAccount = accountsDir.Sub("class") // remove class and attends
	JennyAccount = accountsDir.Sub("attends")

	if err != nil {
		log.Fatalf("Unable to set FDB database value (%v)", err)
	}

	// Data Model for Key: ("AccountBalance", person, balance) = ""

	loadAccount(db, "Tim", 400)
	fetchAccount(db, "Tim", 400)
	// transferMoney(db, "Tim", 200, 200)
	//fetchAccount(db, "Jenny", 200)
	test, _ := listAllAccounts(db)
	fmt.Println("this is test", test)
}

func loadAccount(t fdb.Transactor, person string, amount int) (err error) {
	SCKey := TimAccount.Pack(tuple.Tuple{person, amount})
	fmt.Println(SCKey)
	// print encoding keys, more info.: https://forums.foundationdb.org/t/application-design-using-subspace-and-tuple/452

	// converting int (amount) to bytes to use it in "Set" method
	buf := new(bytes.Buffer)
	error_ := binary.Write(buf, binary.LittleEndian, int32(amount))
	if error_ != nil {
		fmt.Println("binary.Write failed:", error_)
	}
	fmt.Printf("% x", buf.Bytes())

	_, err = t.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {

		tr.Set(SCKey, buf.Bytes()) // we set an encoded key out of the Tuple of person and amount

		return
	})
	return
}

func transferMoney(t fdb.Transactor, source string, target, amount int) (err error) {

	// money transfer consist out of the following methods:
	// - Transact: withdraw money
	// - Transaction: load money on other account
	// - if error -> recover to previous state

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
	fmt.Println(ret)

	v := ret.([]byte)
	fmt.Printf("func fetchAccount called: %s\n", string(v))
	fmt.Println("this is v: ", v)
	data := binary.LittleEndian.Uint32(v) // decoding the byte with the help of binary package
	fmt.Println(data)

	return

}

func listAllAccounts(t fdb.Transactor) (ac interface{}, err error) {
	//var personAccount personalAccount
	var allAccounts []personalAccount
	r, err := t.ReadTransact(func(rtr fdb.ReadTransaction) (interface{}, error) {
		ri := rtr.GetRange(TimAccount, fdb.RangeOptions{}).Iterator()
		for ri.Advance() {
			kv := ri.MustGet()
			t, err := TimAccount.Unpack(kv.Key)
			if err != nil {
				return nil, err
			}
			personAccount := personalAccount{t[0].(string), t[1].(int64)}
			allAccounts = append(allAccounts, personAccount)

		}
		account := accountList{"bankABC", allAccounts}
		fmt.Println(account)

		return account, nil
	})
	if err != nil {
		fmt.Println("called in error: ", r)
		//ac = r.([]string)
	}
	return r, err
}

func dropAccount(t fdb.Transactor, person, amount int) (err error) {
	SCKey := TimAccount.Pack(tuple.Tuple{person, amount})

	_, err = t.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Clear(SCKey)
		return
	})
	return
}
