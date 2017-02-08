package popcount

import (
	"sync"
)

var loadTableOnce sync.Once
var pc [256]byte

func init() {
	// This initialization is fairly simple,
	// so synchronization might not give performance a boost.
	loadTableOnce.Do(func() {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
	})
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var count byte
	for i := uint64(0); i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}
