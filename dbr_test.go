package main

import "testing"

func BenchmarkDbrRead(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		DbrRead(i)
	}
}