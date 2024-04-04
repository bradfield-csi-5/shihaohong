package memtable

import (
	"fmt"
	"testing"
)

// To run, `go test -bench=.`
const MAX_ITEMS = 5000

func BenchmarkMapMemtablePut(b *testing.B) {
	mt := NewMapMT()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			err := mt.Put(key, val)
			if err != nil {
				panic(err)
			}
		}
		mt = NewMapMT()
	}
}

func BenchmarkMapMemtableGetMiddle(b *testing.B) {
	mt := NewMapMT()
	for j := 0; j < MAX_ITEMS; j++ {
		key := []byte("key" + fmt.Sprint(j))
		val := []byte("item" + fmt.Sprint(j))
		mt.Put(key, val)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mt.Get([]byte("key" + fmt.Sprint(MAX_ITEMS/2)))
	}
}

func BenchmarkMapMemtablePutGet(b *testing.B) {
	mt := NewMapMT()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			mt.Put(key, val)
			mt.Get(key)
		}
		mt = NewMapMT()
	}
}

func BenchmarkMapMemtablePutDelete(b *testing.B) {
	mt := NewMapMT()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte(String(5) + fmt.Sprint(i))
		val := []byte("item" + fmt.Sprint(i))
		mt.Put(key, val)
	}
	b.ResetTimer()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte(String(5) + fmt.Sprint(i))
		val := []byte(String(5) + fmt.Sprint(i))

		mt.Put(key, val)
		mt.Delete(key)
	}
}

func BenchmarkMapMemtableRangeScan(b *testing.B) {
	alphabet := []string{}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	mt := NewMapMT()
	for i := 0; i < len(alphabet); i++ {
		for j := 0; j < 100; j++ {
			key := []byte(alphabet[i] + fmt.Sprint(j))
			val := []byte("item" + alphabet[i] + fmt.Sprint(j))
			mt.Put(key, val)
		}
	}
	b.ResetTimer()
	iter, _ := mt.RangeScan([]byte("A50"), []byte("B50"))

	// get all iterator values until exhausted
	for iter.Next() {
		// just to verify results of rangescan
		// fmt.Println(string(iter.Key()))
		// fmt.Println(string(iter.Value()))
	}
}
