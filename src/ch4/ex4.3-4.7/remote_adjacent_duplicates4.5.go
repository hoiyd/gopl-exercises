package main

import (
	"fmt"
)

func removeAdjacentDuplicates(strings []string) []string {
	i := 1
	for _, s := range strings {
		if s != strings[i-1] {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	s := []string{"1", "2", "3", "3", "3", "6", "7", "7"}
	s = removeAdjacentDuplicates(s) //1,2,3,6,7
	fmt.Println(s)
}
