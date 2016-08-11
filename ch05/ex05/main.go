package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	words, images := countWordsAndImages(doc)
	fmt.Printf("words: %d\n", words)
	fmt.Printf("images: %d\n", images)
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return scan(0, 0, n)
}

func scan(words, images int, n *html.Node) (int, int) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return words, images
		}

		if n.Data == "img" {
			images++
		}
	}

	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}

	if n.FirstChild != nil {
		words, images = scan(words, images, n.FirstChild)
	}

	if n.NextSibling != nil {
		words, images = scan(words, images, n.NextSibling)
	}

	return words, images
}
