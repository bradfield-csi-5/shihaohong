package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	// The returned DB instance is safe for concurrent use. Which mean that all
	// DB's methods may be called concurrently from multiple goroutine.
	db, err := leveldb.OpenFile("go_db", nil)
	errCheck(err)
	defer db.Close()

	err = db.Put([]byte("person1"), []byte("shihao"), nil)
	errCheck(err)
	fmt.Println("added item1")
	err = db.Put([]byte("person2"), []byte("jennifer"), nil)
	errCheck(err)
	fmt.Println("added item2")
	err = db.Put([]byte("person3"), []byte("jessica"), nil)
	errCheck(err)
	fmt.Println("added item3")
	err = db.Put([]byte("person4"), []byte("collin"), nil)
	errCheck(err)
	fmt.Println("added item4")
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
