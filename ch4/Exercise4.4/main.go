package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice = rotate(slice, 3)
	fmt.Println(slice)
}

func rotate(slice []int, d int) []int {
	n := len(slice)
	d = d % n
	tmp := make([]int, d)
	copy(tmp, slice[:d])
	copy(slice, slice[d:])
	copy(slice[n-d:], tmp)
	return slice
}
