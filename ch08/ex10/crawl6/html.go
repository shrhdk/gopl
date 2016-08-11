package main

import (
	"golang.org/x/net/html"
	"io"
	"net/url"
)

func extractLinks(base *url.URL, r io.Reader) ([]*url.URL, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []*url.URL

	visitNode := func(n *html.Node) {
		if !(n.Type == html.ElementNode && n.Data == "a") {
			return
		}

		v, ok := getAttr(n, "href")
		if !ok {
			return
		}

		link, err := base.Parse(v)
		if err != nil {
			return
		}

		links = append(links, link)
	}

	forEachNode(doc, visitNode, nil)

	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func getAttr(n *html.Node, name string) (val string, ok bool) {
	for _, a := range n.Attr {
		if a.Key == name {
			return a.Val, true
		}
	}
	return "", false
}
