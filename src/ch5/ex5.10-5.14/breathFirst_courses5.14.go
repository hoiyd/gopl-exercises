package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	var courses []string
	for course, _ := range prereqs {
		courses = append(courses, course)
	}
	breadthFirst(get, courses)
}

func get(s string) []string {
	return prereqs[s]
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(get func(string) []string, list []string) {
	visited := make(map[string]bool)
	for len(list) > 0 {
		items := list
		list = nil
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				fmt.Println(item)
				list = append(list, get(item)...)
			}
		}
	}
}
