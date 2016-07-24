package main

import (
	"fmt"
)

func reverseUTF8BytesInPlace(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func main() {
	// convert bytes slice to a string, then to a runes slice.
	bytes := []byte("Hello, 世界")
	s := string(bytes)
	runes := []rune(s)
	runes = reverseUTF8BytesInPlace(runes)
	fmt.Printf("%s\n", string(runes))
}
