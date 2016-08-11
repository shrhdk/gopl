package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var depth int
	forEachNode(doc,
		func(n *html.Node) {
			switch n.Type {
			case html.ElementNode:
				fmt.Print(sprintIndent(depth))
				fmt.Print(sprintElement(n))
			case html.TextNode:
				s := strings.TrimSpace(n.Data)
				if s != "" {
					fmt.Print(sprintIndent(depth))
					fmt.Printf("%s\n", s)
				}
			case html.CommentNode:
				fmt.Print(sprintIndent(depth))
				fmt.Printf("<!-- %s -->\n", n.Data)
			}

			depth++
		},
		func(n *html.Node) {
			depth--

			if n.Type != html.ElementNode || n.FirstChild == nil {
				return
			}

			fmt.Print(sprintIndent(depth))
			fmt.Printf("</%s>\n", n.Data)
		})
}

func sprintIndent(depth int) string {
	return fmt.Sprintf("%*s", depth*2, "")
}

func sprintElement(n *html.Node) string {
	if n.Type != html.ElementNode {
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
