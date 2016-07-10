package main

import (
	"testing"
)

func BenchmarkPqRead(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		PqRead(i)
	}
}
