package main

import (
	"fmt"
	"testing"
)

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
}
