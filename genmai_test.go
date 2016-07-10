package main

import "testing"

func BenchmarkGenmaiRead(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		GenmaiRead(i)
	}
}