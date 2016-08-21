package main

import (
	"fmt"
	"math/rand"
	"time"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return appendValues(values[:0], root)
}

// Recursively extract values(sorted) from a tree,
// first left subtree, then root, then right subtree(中序遍历)
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// Recursively add value to a tree,
// if smaller it goes to left subtree,
// else it goes to right subtree.
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var str string

	var visit func(tr *tree, identation int)

	visit = func(tr *tree, identation int) {
		if identation == 0 {
			str = str + fmt.Sprintf("%*s%d\n", 2*identation, "", tr.value)
		} else {
			str = str + fmt.Sprintf("%*s- %d\n", 2*identation, "", tr.value)
		}

		if tr.left != nil {
			visit(tr.left, identation+1)
		} else {
			str = str + fmt.Sprintf("%*s- %s\n", 2*(identation+1), "", "leftnil")
		}

		if tr.right != nil {
			visit(tr.right, identation+1)
		} else {
			str = str + fmt.Sprintf("%*s- %s\n", 2*(identation+1), "", "rightnil")
		}
	}
	visit(t, 0)
	return str
}

func main() {
	seed := rand.NewSource(time.Now().Unix())
	r := rand.New(seed)

	data := make([]int, 10)
	for i := range data {
		data[i] = r.Intn(100)
	}
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	fmt.Println(root.String())
}
