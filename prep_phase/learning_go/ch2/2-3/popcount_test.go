package main

import (
	"testing"
)

const val uint64 = 255

/*
go test -bench=PopCount
goos: darwin
goarch: amd64
cpu: VirtualApple @ 2.50GHz
BenchmarkPopCount-8             1000000000               0.3305 ns/op
BenchmarkPopCountLoop-8         300548124                3.941 ns/op
BenchmarkPopCountShift-8        54058618                22.25 ns/op
BenchmarkPopCountClear-8        309934998                3.880 ns/op
*/

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(val)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(val)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(val)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(val)
	}
}
