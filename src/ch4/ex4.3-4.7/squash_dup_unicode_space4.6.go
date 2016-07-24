package main

import (
	"fmt"
	"unicode"
)

func removeDupUnicodeSpace(bytes []byte) []byte {
	out := bytes[:0]
	for i := 0; i < len(bytes); i++ {
		// always keep the first rune, whether it is space or not
		if i > 0 && unicode.IsSpace(rune(bytes[i])) {
			if unicode.IsSpace(rune(bytes[i-1])) {
				continue
			} else {
				out = append(out, ' ')
			}
		} else {
			out = append(out, bytes[i])
		}
	}
	return out
}

func main() {
	bytes := []byte("Hello \r \n    , 世界")
	bytes = removeDupUnicodeSpace(bytes)
	fmt.Printf("%q\n", bytes)
}
