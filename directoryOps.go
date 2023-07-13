package main

/*

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)
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

func dropAccount(t fdb.Transactor, person string, amount int) (err error) {
	SCKey := TimAccount.Pack(tuple.Tuple{person, amount})

	_, err = t.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Clear(SCKey)
		return
	})
	return
}
*/
