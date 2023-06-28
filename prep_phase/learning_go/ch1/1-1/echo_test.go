package main

import (
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2()
	}
}

func BenchmarkJoinEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinEcho()
	}
}
