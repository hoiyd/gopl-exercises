package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [8]byte

func main() {
	for i := uint(0); i < 8; i++ {
		pc[i] = byte(1 << i)
	}
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(sha256DiffCount(c1, c2))
}

func sha256DiffCount(c1 [32]byte, c2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		// count += bitCount(c1[i], c2[i])
		count += popCount(c1[i], c2[i]) // or use popCount
	}
	return count
}

func bitCount(x byte, y byte) int {
	z := x ^ y
	count := 0

	for ; z != 0; z = z >> 1 {
		if z&byte(1) == byte(1) {
			count++
		}
	}
	return count
}

func popCount(x byte, y byte) int {
	count := 0
	for j := 0; j < 8; j++ {
		if x&pc[j] != y&pc[j] {
			count++
		}
	}
	return count
}
