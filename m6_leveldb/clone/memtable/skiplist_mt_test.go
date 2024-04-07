package memtable

import (
	"fmt"
	"testing"
)

func BenchmarkSkipListMTPut(b *testing.B) {
	mt := NewSkipListMT()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			err := mt.Put(key, val)
			if err != nil {
				panic(err)
			}
		}
		mt = NewSkipListMT()
	}
}

func BenchmarkSkipListMTGetMiddle(b *testing.B) {
	mt := NewSkipListMT()
	for j := 0; j < MAX_ITEMS; j++ {
		key := []byte("key" + fmt.Sprint(j))
		val := []byte("item" + fmt.Sprint(j))
		err := mt.Put(key, val)
		if err != nil {
			panic(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := mt.Get([]byte("key" + fmt.Sprint(MAX_ITEMS/2)))
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSkipListMTPutGet(b *testing.B) {
	mt := NewSkipListMT()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX_ITEMS; j++ {
			key := []byte(String(5) + fmt.Sprint(j))
			val := []byte("item" + fmt.Sprint(j))
			err := mt.Put(key, val)
			if err != nil {
				panic(err)
			}
			_, err = mt.Get(key)
			if err != nil {
				panic(err)
			}
		}
		mt = NewSkipListMT()
	}
}

func BenchmarkSkipListMTPutDelete(b *testing.B) {
	mt := NewSkipListMT()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte(String(5) + fmt.Sprint(i))
		val := []byte("item" + fmt.Sprint(i))
		err := mt.Put(key, val)
		if err != nil {
			panic(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < MAX_ITEMS; i++ {
		key := []byte(String(5) + fmt.Sprint(i))
		val := []byte(String(5) + fmt.Sprint(i))

		err := mt.Put(key, val)
		if err != nil {
			panic(err)
		}
		err = mt.Delete(key)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSkipListMTRangeScan(b *testing.B) {
	alphabet := []string{}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	mt := NewSkipListMT()
	for i := 0; i < len(alphabet); i++ {
		for j := 0; j < 100; j++ {
			key := []byte(alphabet[i] + fmt.Sprintf("%03d", j))
			val := []byte("item" + alphabet[i] + fmt.Sprintf("%03d", j))
			mt.Put(key, val)
		}
	}
	b.ResetTimer()
	iter, _ := mt.RangeScan([]byte("A050"), []byte("B050"))

	// get all iterator values until exhausted
	for iter.Next() {
		// just to verify results of rangescan
		// fmt.Println(string(iter.Key()))
		// fmt.Println(string(iter.Value()))
	}
}
