package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	found := ElementByID(doc, os.Args[1])

	fmt.Println("found: " + sprintElement(found))
}

func sprintIndent(depth int) string {
	return fmt.Sprintf("%*s", depth*2, "")
}

func sprintElement(n *html.Node) string {
	if n == nil || n.Type != html.ElementNode {
		return ""
	}

	var tag bytes.Buffer

	tag.WriteRune('<')
	tag.WriteString(n.Data)

	if len(n.Attr) > 0 {
		for _, a := range n.Attr {
			if tag.Len() != 0 {
				tag.WriteRune(' ')
			}

			tag.WriteString(a.Key)
			tag.WriteString("='")
			tag.WriteString(a.Val)
			tag.WriteString("'")
		}
	}

	if n.FirstChild != nil {
		tag.WriteString(">\n")
	} else {
		tag.WriteString("/>\n")
	}

	return tag.String()
}

func getAttr(n *html.Node, name string) string {
	for _, a := range n.Attr {
		if a.Key == name {
			return a.Val
		}
	}

	return ""
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var found *html.Node
	var depth int
	forEachNode(doc,
		func(n *html.Node) bool {
			if n.Type != html.ElementNode {
				return true
			}

			if getAttr(n, "id") == id {
				found = n
				return false
			}

			fmt.Print(sprintIndent(depth))
			fmt.Printf(sprintElement(n))

			depth++

			return true
		},
		func(n *html.Node) bool {
			if n.Type != html.ElementNode {
				return true
			}

			depth--

			fmt.Print(sprintIndent(depth))
			fmt.Printf("</%s>\n", n.Data)

			return true
		})

	return found
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil && !pre(n) {
		return false
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil && !post(n) {
		return false
	}

	return true
}
