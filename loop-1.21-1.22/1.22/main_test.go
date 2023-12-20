package main

import (
	"testing"
)

func BenchmarkForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loop()
	}
}
