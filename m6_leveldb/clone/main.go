package main

import "fmt"

func main() {
	db := NewMemoryDB()
	err := db.Put([]byte("Aperson1"), []byte("shihao"))
	errCheck(err)
	fmt.Println("added item1")
	err = db.Put([]byte("Dperson2"), []byte("jennifer"))
	errCheck(err)
	fmt.Println("added item2")
	err = db.Put([]byte("Zperson3"), []byte("jessica"))
	errCheck(err)
	fmt.Println("added item3")
	err = db.Put([]byte("Hperson4"), []byte("collin"))
	errCheck(err)
	fmt.Println("added item4")
	val, err := db.Get([]byte("Dperson2"))
	errCheck(err)
	fmt.Printf("Dperson2 val: %s\n", val)
	err = db.Delete([]byte("Dperson2"))
	errCheck(err)
	fmt.Println("deleted Dperson2")
	val, err = db.Get([]byte("Dperson2"))
	errCheck(err)
	fmt.Printf("Dperson2 val: %s\n", val)
	err = db.Put([]byte("Dperson2"), []byte("jennifer"))
	errCheck(err)
	val, err = db.Get([]byte("Dperson2"))
	errCheck(err)
	fmt.Printf("Dperson2 val: %s\n", val)

	// valid values
	it, err := db.RangeScan([]byte("Dperson2"), []byte("Zperson3"))
	errCheck(err)
	iterKey := it.Key()
	iterVal := it.Value()
	fmt.Printf("iterKey: %s\n", iterKey)
	fmt.Printf("iterVal: %s\n", iterVal)
	ok := it.Next()
	fmt.Printf("hasNewVal: %t\n", ok)
	ok = it.Next()
	fmt.Printf("hasNewVal: %t\n", ok)
	iterKey = it.Key()
	iterVal = it.Value()
	fmt.Printf("iterKey: %s\n", iterKey)
	fmt.Printf("iterVal: %s\n", iterVal)
	ok = it.Next()
	fmt.Printf("hasNewVal: %t\n", ok)

	// invalid values
	_, err = db.RangeScan([]byte("1"), []byte("2"))
	iterKey = it.Key()
	iterVal = it.Value()
	fmt.Printf("iterKey: %s\n", iterKey)
	fmt.Printf("iterVal: %s\n", iterVal)
	errCheck(err)
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
