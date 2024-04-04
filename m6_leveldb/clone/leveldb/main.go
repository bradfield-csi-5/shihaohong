package main

import "fmt"

func main() {
	db := NewLevelDB("wal.01")
	db.Init()

	// log.ClearLog()
	// log.Put([]byte("chris"), []byte("chris' item"))
	// log.Put([]byte("shi hao"), []byte("shi hao's item"))
	// log.Put([]byte("luke"), []byte("luke's item"))
	// log.Put([]byte("ben"), []byte("ben's item"))
	// log.Delete([]byte("shi hao"))
	// entries, err := log.Read()
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(entries); i++ {
	// 	fmt.Println(entries[i])
	// }

	// mt := memtable.NewSkipListMT()

	fmt.Println("nodes")
	res, _ := db.Get([]byte("luke"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("ben"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("chris"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("shi hao"))
	fmt.Println(string(res))
}

// func main() {
// 	mt := NewMemorymt()
// 	err := mt.Put([]byte("Aperson1"), []byte("shihao"))
// 	errCheck(err)
// 	fmt.Println("added item1")
// 	err = mt.Put([]byte("Dperson2"), []byte("jennifer"))
// 	errCheck(err)
// 	fmt.Println("added item2")
// 	err = mt.Put([]byte("Zperson3"), []byte("jessica"))
// 	errCheck(err)
// 	fmt.Println("added item3")
// 	err = mt.Put([]byte("Hperson4"), []byte("collin"))
// 	errCheck(err)
// 	fmt.Println("added item4")
// 	val, err := mt.Get([]byte("Dperson2"))
// 	errCheck(err)
// 	fmt.Printf("Dperson2 val: %s\n", val)
// 	err = mt.Delete([]byte("Dperson2"))
// 	errCheck(err)
// 	fmt.Println("deleted Dperson2")
// 	val, err = mt.Get([]byte("Dperson2"))
// 	errCheck(err)
// 	fmt.Printf("Dperson2 val: %s\n", val)
// 	err = mt.Put([]byte("Dperson2"), []byte("jennifer"))
// 	errCheck(err)
// 	val, err = mt.Get([]byte("Dperson2"))
// 	errCheck(err)
// 	fmt.Printf("Dperson2 val: %s\n", val)

// 	// valid values
// 	it, err := mt.RangeScan([]byte("Dperson2"), []byte("Zperson3"))
// 	errCheck(err)
// 	iterKey := it.Key()
// 	iterVal := it.Value()
// 	fmt.Printf("iterKey: %s\n", iterKey)
// 	fmt.Printf("iterVal: %s\n", iterVal)
// 	ok := it.Next()
// 	fmt.Printf("hasNewVal: %t\n", ok)
// 	ok = it.Next()
// 	fmt.Printf("hasNewVal: %t\n", ok)
// 	iterKey = it.Key()
// 	iterVal = it.Value()
// 	fmt.Printf("iterKey: %s\n", iterKey)
// 	fmt.Printf("iterVal: %s\n", iterVal)
// 	ok = it.Next()
// 	fmt.Printf("hasNewVal: %t\n", ok)

// 	// invalid values
// 	_, err = mt.RangeScan([]byte("1"), []byte("2"))
// 	iterKey = it.Key()
// 	iterVal = it.Value()
// 	fmt.Printf("iterKey: %s\n", iterKey)
// 	fmt.Printf("iterVal: %s\n", iterVal)
// 	errCheck(err)
// }

// func errCheck(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
