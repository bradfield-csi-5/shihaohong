package main

import (
	"fmt"
	"testing"
)

func BenchmarkArrDBPut(b *testing.B) {
	db := NewArrDB()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			db.Put(key, val)
		}
		db = NewArrDB()
	}
}

func BenchmarkArrDBGetMiddle(b *testing.B) {
	db := NewArrDB()
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

func BenchmarkArrDBPutGet(b *testing.B) {
	db := NewArrDB()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			db.Put(key, val)
			db.Get(key)
		}
		db = NewArrDB()
	}
}
