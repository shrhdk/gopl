package ex07

import "golang.org/x/net/html"

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var elems []*html.Node
	forEachNode(doc, nil, func(n *html.Node) {
		if n.Type != html.ElementNode {
			return
		}

		if !contains(name, n.Data) {
			return
		}

		elems = append(elems, n)
	})
	return elems
}

func contains(list []string, item string) bool {
	for _, s := range list {
		if s == item {
			return true
		}
	}

	return false
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
