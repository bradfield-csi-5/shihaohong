package main

import (
	"fmt"
	"testing"
)

// To run, `go test -bench=.`
const MAX_ITEMS = 100000

func BenchmarkMemoryDBPut(b *testing.B) {
	db := NewMemoryDB()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte("key" + fmt.Sprint(i))
		val := []byte("item" + fmt.Sprint(i))
		db.Put(key, val)
	}
}

func BenchmarkMemoryDBGet(b *testing.B) {
	db := NewMemoryDB()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte("key" + fmt.Sprint(i))
		val := []byte("item" + fmt.Sprint(i))
		db.Put(key, val)
	}
	b.ResetTimer()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte("key" + fmt.Sprint(i))
		db.Get(key)
	}
}

func BenchmarkMemoryDBDelete(b *testing.B) {
	db := NewMemoryDB()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte("key" + fmt.Sprint(i))
		val := []byte("item" + fmt.Sprint(i))
		db.Put(key, val)
	}
	b.ResetTimer()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte("key" + fmt.Sprint(i))
		db.Delete(key)
	}
}

func BenchmarkMemoryDBRangeScan(b *testing.B) {
	alphabet := []string{}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	db := NewMemoryDB()
	for i := 0; i < len(alphabet); i++ {
		for j := 0; j < 100; j++ {
			key := []byte(alphabet[i] + fmt.Sprint(j))
			val := []byte("item" + alphabet[i] + fmt.Sprint(j))
			db.Put(key, val)
		}
	}
	b.ResetTimer()
	iter, _ := db.RangeScan([]byte("A50"), []byte("B50"))

	// get all iterator values until exhausted
	for iter.Next() {
		// just to verify results of rangescan
		// fmt.Println(string(iter.Key()))
		// fmt.Println(string(iter.Value()))
	}
}
