package main

import (
	"testing"
)

func BenchmarkGorpRead(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		GorpRead(i)
	}
}
