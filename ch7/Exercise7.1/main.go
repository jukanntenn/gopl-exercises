package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	tmp := 0
	for scanner.Scan() {
		tmp++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*c += WordCounter(tmp)
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	tmp := 0
	for scanner.Scan() {
		tmp++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*c += LineCounter(tmp)
	return len(p), nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"))
	fmt.Println(wc)

	wc = 0
	var name = "Dolly"
	fmt.Fprintf(&wc, "hello, %s", name)
	fmt.Println(wc)

	var lc LineCounter
	lc.Write([]byte("Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"))
	fmt.Println(lc)

	lc = 0
	fmt.Fprintf(&lc, "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n")
	fmt.Println(lc)
}
