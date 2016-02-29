package main

import (
	"testing"

	pc "ch2/ex2.3-5/popcount"
)

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopcountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.PopCountLoop(0x1234567890ABCDEF)
	}
}

// BenchmarkPopcount-8    	200000000	         6.02 ns/op
// BenchmarkPopcountLoop-8	100000000	        15.0 ns/op
