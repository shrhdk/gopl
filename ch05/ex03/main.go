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
	for _, line := range visit(nil, doc) {
		fmt.Println(line)
	}
}

func visit(text []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return text
		}
	}

	if n.Type == html.TextNode {
		t := strings.TrimSpace(n.Data)
		if t != "" {
			text = append(text, t)
		}
	}

	if n.FirstChild != nil {
		text = visit(text, n.FirstChild)
	}

	if n.NextSibling != nil {
		text = visit(text, n.NextSibling)
	}

	return text
}
