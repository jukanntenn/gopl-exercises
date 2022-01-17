package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq := make(map[string]int)
	f, err := os.Open("words.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		wordfreq[input.Text()]++
	}
	fmt.Println("word\tfreq")
	for k, v := range wordfreq {
		fmt.Printf("%s  %d\n", k, v)
	}
}
