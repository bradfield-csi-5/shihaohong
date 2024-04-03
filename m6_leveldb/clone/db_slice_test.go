package main

import (
	"fmt"
	"testing"
)

func BenchmarkSliceDBPut(b *testing.B) {
	db := NewSliceDB()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			err := db.Put(key, val)
			if err != nil {
				panic(err)
			}
		}
		db = NewSliceDB()
	}
}

func BenchmarkSliceDBGetMiddle(b *testing.B) {
	db := NewSliceDB()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte("key" + fmt.Sprint(i))
		val := []byte("item" + fmt.Sprint(i))
		err := db.Put(key, val)
		if err != nil {
			panic(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := db.Get([]byte("key" + fmt.Sprint(MAX_ITEMS/2)))
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSliceDBPutGet(b *testing.B) {
	db := NewSliceDB()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			err := db.Put(key, val)
			if err != nil {
				panic(err)
			}
			_, err = db.Get(key)
			if err != nil {
				panic(err)
			}
		}
		db = NewSliceDB()
	}
}

func BenchmarkSliceDBPutDelete(b *testing.B) {
	db := NewSliceDB()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte(String(5) + fmt.Sprint(i))
		val := []byte("item" + fmt.Sprint(i))
		err := db.Put(key, val)
		if err != nil {
			panic(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte(String(5) + fmt.Sprint(i))
		val := []byte(String(5) + fmt.Sprint(i))

		err := db.Put(key, val)
		if err != nil {
			panic(err)
		}
		err = db.Delete(key)
		if err != nil {
			panic(err)
		}
	}
}
