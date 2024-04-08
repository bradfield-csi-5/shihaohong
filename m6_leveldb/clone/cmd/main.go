package main

import (
	"fmt"

	"github.com/shihaohong/leveldb_clone/leveldb"
)

func main() {
	db := leveldb.NewLevelDB("wal.01", "sst")
	// db.Init()

	// db.ClearWAL()
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
	fmt.Println("put:")
	db.Put([]byte("another key"), []byte("new val"))
	db.Put([]byte("shihao key"), []byte("sh"))
	db.Put([]byte("chris key"), []byte("ch"))
	db.Put([]byte("luke key"), []byte("lu"))
	// db.Delete([]byte("another key"))
	// db.Delete([]byte("shihao key"))
	// db.Delete([]byte("chris key"))
	// db.Delete([]byte("luke key"))

	fmt.Println("nodes")
	res, _ := db.Get([]byte("another key"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("shihao key"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("chris key"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("luke key"))
	fmt.Println(string(res))
	db.Store()
	fmt.Println("nodes")
	res, _ = db.Get([]byte("another key"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("shihao key"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("chris key"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("luke key"))
	fmt.Println(string(res))

	db.Put([]byte("another key"), []byte("new val"))
	db.Put([]byte("shihao key"), []byte("sh"))
	db.Put([]byte("chris key"), []byte("ch"))
	db.Put([]byte("luke key"), []byte("lu"))
	res, _ = db.Get([]byte("another key"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("shihao key"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("chris key"))
	fmt.Println(string(res))
	res, _ = db.Get([]byte("luke key"))
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
