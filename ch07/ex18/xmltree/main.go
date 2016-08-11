package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	tree, err := parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(tree)
}

func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("xmlselect: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) != 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, elem)
			}
			stack = append(stack, elem) // push
		case xml.EndElement:
			if len(stack) == 1 {
				break
			}
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if strings.TrimSpace(string(tok)) == "" {
				continue
			}

			elem := CharData(tok)
			if len(stack) >= 1 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, elem)
			}
		}
	}
	return stack[0], nil
}

type Node interface {
	String() string
}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (c CharData) String() string {
	return string(c)
}

func (e Element) String() string {
	var buf bytes.Buffer
	if len(e.Attr) == 0 {
		buf.WriteString(fmt.Sprintf("<%s>\n", e.Type.Local))
	} else {
		buf.WriteString(fmt.Sprintf("<%s", e.Type.Local))
		for _, a := range e.Attr {
			buf.WriteString(fmt.Sprintf(" %s=\"%s\"", a.Name.Local, a.Value))
		}
		buf.WriteString(">\n")
	}
	for _, n := range e.Children {
		buf.WriteString(n.String())
	}
	buf.WriteString(fmt.Sprintf("</%s>\n", e.Type.Local))
	return buf.String()
}
