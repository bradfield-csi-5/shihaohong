package main

import (
	"fmt"
	"testing"
)

func BenchmarkSkipListDBPut(b *testing.B) {
	db := NewSkipListDB()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte("key" + fmt.Sprint(i))
		val := []byte("item" + fmt.Sprint(i))
		db.Put(key, val)
	}
}

func BenchmarkSkipListDBGet(b *testing.B) {
	db := NewSkipListDB()
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
