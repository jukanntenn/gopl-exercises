package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages("https://www.studygolang.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v", err)
		os.Exit(1)
	}
	fmt.Printf("wordss: %d\nimages: %d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		splited := strings.Split(n.Data, " ")
		words += len(splited)
	}
	if n.FirstChild != nil {
		w, i := countWordsAndImages(n.FirstChild)
		words += w
		images += i
	}
	if n.NextSibling != nil {
		w, i := countWordsAndImages(n.NextSibling)
		words += w
		images += i
	}
	return
}
