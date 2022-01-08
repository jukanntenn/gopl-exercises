package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]map[string]struct{})
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, names)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			keys := make([]string, len(names[line]))
			for k := range names[line] {
				keys = append(keys, k)
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(keys, " "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, names map[string]map[string]struct{}) {
	input := bufio.NewScanner(f)
	// for inputs from stdin, press Ctrl + D to end the loop
	for input.Scan() {
		text := input.Text()
		if counts[text] == 0 {
			names[text] = make(map[string]struct{})
		}
		counts[text]++
		names[text][f.Name()] = struct{}{}
	}
	// NOTE: ignoring potential errors from input.Err()
}
