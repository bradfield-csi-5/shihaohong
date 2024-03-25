package main

import "fmt"

func main() {
	db := NewMemoryDB()
	err := db.Put([]byte("person1"), []byte("shihao"))
	errCheck(err)
	fmt.Println("added item1")
	err = db.Put([]byte("person2"), []byte("jennifer"))
	errCheck(err)
	fmt.Println("added item2")
	err = db.Put([]byte("person3"), []byte("jessica"))
	errCheck(err)
	fmt.Println("added item3")
	err = db.Put([]byte("person4"), []byte("collin"))
	errCheck(err)
	fmt.Println("added item4")
	val, err := db.Get([]byte("person2"))
	errCheck(err)
	fmt.Printf("person2 val: %s\n", val)
	err = db.Delete([]byte("person2"))
	errCheck(err)
	fmt.Println("deleted person2")
	val, err = db.Get([]byte("person2"))
	errCheck(err)
	fmt.Printf("person2 val: %s\n", val)
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
