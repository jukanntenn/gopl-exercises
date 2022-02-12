package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exercise5.2: %v\n", err)
		os.Exit(1)
	}
	counter := count_tags(make(map[string]int), doc)
	for k, v := range counter {
		fmt.Printf("%d %s\n", v, k)
	}
}

func count_tags(counter map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}
	if n.FirstChild != nil {
		count_tags(counter, n.FirstChild)
	}
	if n.NextSibling != nil {
		count_tags(counter, n.NextSibling)
	}
	return counter
}
