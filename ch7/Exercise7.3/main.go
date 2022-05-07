package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
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
	var s string
	if t == nil {
		return s
	}

	s += t.left.String()
	s += strconv.Itoa(t.value)
	s += t.right.String()

	return s
}

func main() {
	var t *tree
	t = add(t, 5) // get root of tree
	add(t, 1)
	add(t, 4)
	add(t, 7)
	add(t, 6)
	add(t, 3)
	add(t, 8)
	add(t, 2)
	fmt.Println(t.String())
}
