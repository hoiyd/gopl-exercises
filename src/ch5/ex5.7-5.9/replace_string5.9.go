package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "foo$bar"
	s2 := expand(s, "bar", func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println(s2)
}

func expand(s, sub string, f func(string) string) string {
	return strings.Replace(s, "$"+sub, f(sub), -1)
}
