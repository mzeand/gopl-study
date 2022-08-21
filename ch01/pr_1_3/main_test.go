package main

import (
	"testing"
)

func BenchmarkAppend_naive(b *testing.B) {
	args := []string{
		"Apple", "Orange", "Banana", "Strawberry",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		naive(args)
	}
}

func BenchmarkAppend_join(b *testing.B) {
	args := []string{
		"Apple", "Orange", "Banana", "Strawberry",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		join(args)
	}
}
