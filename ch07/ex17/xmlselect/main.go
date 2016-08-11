package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
)

// Examples:
// curl -s https://www.w3.org/TR/2006/REC-xml11-20060816/ | ./xmlselect html body div.head h2
// curl -s https://www.w3.org/TR/2006/REC-xml11-20060816/ | ./xmlselect html body div h2

func main() {
	var path []xml.StartElement
	for _, v := range os.Args[1:] {
		if elem, err := parse(v); err == nil {
			path = append(path, elem)
		}
	}

	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, path) {
				fmt.Printf("%s: %s\n", os.Args[1:], tok)
			}
		}
	}
}

// Parse elem#id elem.class elem class1.class2
func parse(s string) (xml.StartElement, error) {
	nid := strings.Count(s, "#")
	nclass := strings.Count(s, ".")

	if nid >= 2 {
		return xml.StartElement{}, errors.New("multiple id is not allowed")
	}

	if nid == 1 && nclass >= 1 {
		return xml.StartElement{}, errors.New("id and class are exclusive.")
	}

	var elem xml.StartElement

	if nid == 1 {
		ss := strings.Split(s, "#")
		elem.Name = xml.Name{"", ss[0]}
		elem.Attr = append(elem.Attr, xml.Attr{xml.Name{"", "id"}, ss[1]})
		return elem, nil
	}

	if nclass >= 1 {
		ss := strings.Split(s, ".")
		elem.Name = xml.Name{"", ss[0]}
		elem.Attr = append(elem.Attr, xml.Attr{xml.Name{"", "class"}, strings.Join(ss[1:], " ")})
		return elem, nil
	}

	elem.Name = xml.Name{"", s}
	return elem, nil
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []xml.StartElement) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if elemEquals(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func elemEquals(x, y xml.StartElement) bool {
	if x.Name.Local != y.Name.Local {
		return false
	}

	if yid, ok := getAttr(y.Attr, "id"); ok {
		if xid, ok := getAttr(x.Attr, "id"); ok {
			return xid == yid
		}

		return false
	}

	if yc, ok := getAttr(y.Attr, "class"); ok {
		if xc, ok := getAttr(x.Attr, "class"); ok {
			xca := strings.Fields(xc)
			sort.Strings(xca)
			yca := strings.Fields(yc)
			sort.Strings(yca)
			return reflect.DeepEqual(xca, yca)
		}
		return false
	}

	return true
}

func getAttr(attr []xml.Attr, name string) (value string, ok bool) {
	for _, a := range attr {
		if a.Name.Local == name {
			return a.Value, true
		}
	}
	return "", false
}
