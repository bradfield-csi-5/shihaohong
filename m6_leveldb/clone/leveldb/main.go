package main

import (
	"fmt"
	"github.com/shihaohong/leveldb_clone/memtable"
	"github.com/shihaohong/leveldb_clone/wal"
)

func main() {
	log := wal.NewLog("wal.01")
	log.ClearLog()
	log.Put([]byte("item1"), []byte("value1"))
	log.Put([]byte("item2"), []byte("value2"))
	log.Put([]byte("item3"), []byte("value3"))
	log.Put([]byte("item4"), []byte("value4"))
	log.Put([]byte("item5"), []byte("value5"))
	log.Delete([]byte("item2"))

	mt := memtable.NewSkipListMT()

	mt.Put([]byte("chris"), []byte("chris' item"))
	mt.Put([]byte("shi hao"), []byte("shi hao's item"))
	mt.Put([]byte("ben"), []byte("ben's item"))
	mt.Put([]byte("luke"), []byte("luke's item"))

	fmt.Println("nodes")
	res, _ := mt.Get([]byte("luke"))
	fmt.Println(string(res))
	res, _ = mt.Get([]byte("ben"))
	fmt.Println(string(res))
	res, _ = mt.Get([]byte("chris"))
	fmt.Println(string(res))
	res, _ = mt.Get([]byte("shi hao"))
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
