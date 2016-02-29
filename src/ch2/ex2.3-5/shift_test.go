package main

import (
	"testing"
)

func shift(x uint64) int {
	var count int
	for i := uint64(0); i < 64; i++ {
		if (x>>i)&1 == 1 {
			count++
		}
	}
	return count
}

func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shift(0x1234567890ABCDEF)
	}
}

// BenchmarkPopcount-8    	200000000	         6.44 ns/op
// BenchmarkPopcountLoop-8	100000000	        16.4 ns/op
// BenchmarkShift-8       	20000000	        95.1 ns/op
