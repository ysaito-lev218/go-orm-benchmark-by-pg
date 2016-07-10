package main

import (
	"testing"
)

func BenchmarkXormRead(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		XormRead(i)
	}
}
