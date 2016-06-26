package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(isAnagram("level", "eellv"))
	fmt.Println(isAnagram("apple", "elapp"))
	fmt.Println(isAnagram("anagram", "mraaang"))
	fmt.Println(isAnagram("boomerang", "gboonamre"))
	fmt.Println(isAnagram("elephant", "tanleeph"))
}

func isAnagram(a string, b string) bool {
	if a == b || len(a) != len(b) {
		return false
	}

	return sortString(a) == sortString(b)
}

func sortString(s string) string {
	t := strings.Split(s, "")
	sort.Strings(t)
	return strings.Join(t, "")
}
