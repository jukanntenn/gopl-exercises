package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exercise5.3: %v\n", err)
		os.Exit(1)
	}
	print_text_node_content(doc)
}

func print_text_node_content(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}

	if n.Type == html.TextNode && n.Data != "" {
		fmt.Println(n.Data)
	}

	if n.FirstChild != nil {
		print_text_node_content(n.FirstChild)
	}
	if n.NextSibling != nil {
		print_text_node_content(n.NextSibling)
	}
}
