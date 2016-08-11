package main

import (
	"fmt"
	"os"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming", "networks"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSortWithRingDetection(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	//Ring detected: data structures -> discrete math -> networks -> operating systems -> data structures
}

func indexOf(s string, slice []string) int {
	for i, v := range slice {
		if s == v {
			return i
		}
	}
	return -1
}

func topoSortWithRingDetection(m map[string][]string) []string {
	var order, stack []string
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if i := indexOf(item, stack); i != -1 {
				fmt.Printf("Ring detected: %s\n", strings.Join(append(stack[i:], item), " -> "))
				os.Exit(1)
			}
			if !seen[item] {
				seen[item] = true
				stack = append(stack, item) // stack "push"
				visitAll(m[item])
				order = append(order, item)
				stack = stack[:len(stack)-1] // stack "pop"
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}
