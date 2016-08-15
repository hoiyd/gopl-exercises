package main

import (
	"fmt"
	"math"
	// "os"
)

func main() {
	max, _ := max(2, 3, 4, 5, 6, 7, 645, 6, 32)
	min, _ := min(2, 3, 4, 5, 6, 7, 645, 6, 32)
	fmt.Printf("max is %d, min is %d\n", max, min)
}

func max(vals ...int) (result int, ok bool) {
	result = math.MinInt64
	if len(vals) == 0 {
		return
	} else {
		for _, val := range vals {
			if val > result {
				result = val
			}
		}
	}
	return
}

func min(vals ...int) (result int, ok bool) {
	result = math.MaxInt64
	if len(vals) == 0 {
		return
	} else {
		for _, val := range vals {
			if val < result {
				result = val
			}
		}
	}
	return
}
