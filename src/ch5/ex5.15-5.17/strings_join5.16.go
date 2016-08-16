package main

import (
	"fmt"
)

func main() {
	strings := []string{"1", "2", "3", "4", "5"}
	s := Join("=>", strings...)
	fmt.Println(s)
}

// Variadic parameters must be the last function parameter
func Join(delimiter string, vals ...string) string {
	var s string
	length := len(vals)
	for _, val := range vals[:length-1] {
		s = s + val + delimiter
	}
	s = s + vals[length-1]
	return s
}
