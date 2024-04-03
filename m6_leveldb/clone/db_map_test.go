package main

import (
	"fmt"
	"testing"
)

// To run, `go test -bench=.`
const MAX_ITEMS = 5000

func BenchmarkMapDBPut(b *testing.B) {
	db := NewMapDB()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			err := db.Put(key, val)
			if err != nil {
				panic(err)
			}
		}
		db = NewMapDB()
	}
}

func BenchmarkMapDBGetMiddle(b *testing.B) {
	db := NewMapDB()
	for j := 0; j < MAX_ITEMS; j++ {
		key := []byte("key" + fmt.Sprint(j))
		val := []byte("item" + fmt.Sprint(j))
		db.Put(key, val)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Get([]byte("key" + fmt.Sprint(MAX_ITEMS/2)))
	}
}

func BenchmarkMapDBPutGet(b *testing.B) {
	db := NewMapDB()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			db.Put(key, val)
			db.Get(key)
		}
		db = NewMapDB()
	}
}

func BenchmarkMapDBPutDelete(b *testing.B) {
	db := NewMapDB()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte(String(5) + fmt.Sprint(i))
		val := []byte("item" + fmt.Sprint(i))
		db.Put(key, val)
	}
	b.ResetTimer()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte(String(5) + fmt.Sprint(i))
		val := []byte(String(5) + fmt.Sprint(i))

		db.Put(key, val)
		db.Delete(key)
	}
}

func BenchmarkMapDBRangeScan(b *testing.B) {
	alphabet := []string{}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	db := NewMapDB()
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
