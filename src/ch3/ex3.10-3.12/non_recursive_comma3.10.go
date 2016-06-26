package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(noRecursiveComma("12"))
	fmt.Println(noRecursiveComma("123"))
	fmt.Println(noRecursiveComma("1234"))
	fmt.Println(noRecursiveComma("12345"))
	fmt.Println(noRecursiveComma("123456"))
	fmt.Println(noRecursiveComma("1234567"))
}

func noRecursiveComma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	r := n % 3
	var buf bytes.Buffer
	if r > 0 {
		buf.WriteString(s[:r])
		buf.WriteString(",")
	}

	for i := r; i < n; i = i + 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < n {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
