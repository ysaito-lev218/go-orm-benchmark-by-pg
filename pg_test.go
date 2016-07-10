package main

import (
	"testing"
)

func BenchmarkPgRead(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		PgRead(i)
	}
}
