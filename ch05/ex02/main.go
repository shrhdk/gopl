package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for elem, num := range visit(make(map[string]int), doc) {
		fmt.Printf("%s\t%d\n", elem, num)
	}
}

func visit(elems map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return elems
	}

	if n.Type == html.ElementNode {
		elems[n.Data]++
	}

	elems = visit(elems, n.FirstChild)
	elems = visit(elems, n.NextSibling)

	return elems
}
