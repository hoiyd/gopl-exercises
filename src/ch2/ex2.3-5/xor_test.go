package main

import (
	"testing"
)

func xor(x uint64) int {
	var count int
	for ; x != 0; x = x & (x - 1) {
		count++
	}
	return count
}

func BenchmarkXorShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shift(0x1234567890ABCDEF)
	}
}

// BenchmarkPopcount-8    	200000000	         6.87 ns/op
// BenchmarkPopcountLoop-8	100000000	        16.1 ns/op
// BenchmarkShift-8       	20000000	        89.3 ns/op
// BenchmarkXorShift-8    	20000000	        86.8 ns/op
