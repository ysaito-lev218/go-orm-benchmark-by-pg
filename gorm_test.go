package main

import (
	"testing"
)

func BenchmarkGormRead(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		GormRead(i)
	}
}
