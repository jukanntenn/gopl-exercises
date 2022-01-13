package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("%d\n", KB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%d\n", TB)
	fmt.Printf("%d\n", PB)
	fmt.Printf("%d\n", EB)
	// overflow
	// fmt.Printf("%d\n", ZB)
	// fmt.Printf("%d\n", YB)
}
