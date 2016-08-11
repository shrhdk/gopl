package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			links = append(links, getAttr(n.Attr, "href"))
		case "img":
			links = append(links, getAttr(n.Attr, "src"))
		case "script":
			links = append(links, getAttr(n.Attr, "src"))
		case "link":
			links = append(links, getAttr(n.Attr, "href"))
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

func getAttr(attrs []html.Attribute, name string) string {
	for _, a := range attrs {
		if a.Key == name {
			return a.Val
		}
	}
	return ""
}
