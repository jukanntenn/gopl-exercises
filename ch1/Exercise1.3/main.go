package main

import (
	"fmt"
	"time"
)

/*
use `go run . arg1 arg2 arg3` instead of `go run main.go arg1 arg2 arg3`,
otherwise you will got:

	# command-line-arguments
	./main.go:10:2: undefined: echo2
	./main.go:15:2: undefined: echo3

See: https://stackoverflow.com/questions/28081486/how-can-i-go-run-a-project-with-multiple-files-in-the-main-package
*/
func main() {
	echo2()
	echo3()

	start := time.Now()
	echo2()
	fmt.Printf("%dns elapsed\n", time.Since(start).Nanoseconds())
	// go run . arg1 arg2 arg3 arg4 arg5 arg6 arg7
	// output: 2510ns elapsed

	start = time.Now()
	echo3()
	fmt.Printf("%dns elapsed\n", time.Since(start).Nanoseconds())
	// go run . arg1 arg2 arg3 arg4 arg5 arg6 arg7
	// output: 1655ns elapsed
}
