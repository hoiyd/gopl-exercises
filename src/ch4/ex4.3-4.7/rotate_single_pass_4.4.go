package main

import (
	"fmt"
)

func rotate(s []int, n int, right bool) {
	l := len(s)
	tmp := make([]int, n)

	if right {
		copy(tmp, s[(l-n):])
		copy(s[n:], s)
		copy(s, tmp)
	} else {
		copy(tmp, s[:n])
		copy(s, s[n:])
		copy(s[n:], tmp)
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate(s, 2, true) // rotate 2 right. s should be: 7,8,1,2,3,4,5,6
	fmt.Println(s)
	rotate(s, 4, false) // then rotate 4 left. s should be: 3,4,5,6,7,8,1,2
	fmt.Println(s)
}
