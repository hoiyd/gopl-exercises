package main

import (
	"fmt"
)

type PalindromeChecker []byte

func (x PalindromeChecker) Len() int           { return len(x) }
func (x PalindromeChecker) Less(i, j int) bool { return x[i] < x[j] }
func (x PalindromeChecker) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func isPalindrome(x PalindromeChecker) bool {
	for i, j := 0, x.Len()-1; i < j; i, j = i+1, j-1 {
		if !x.Less(i, j) && !x.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(PalindromeChecker([]byte("level"))))
	fmt.Println(isPalindrome(PalindromeChecker([]byte("abcdedcba"))))
}
