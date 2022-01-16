package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(ptr *[10]int) {
	for i, j := 0, 9; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}
