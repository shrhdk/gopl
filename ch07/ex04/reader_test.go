package main

import (
	"golang.org/x/net/html"
	"os"
)

func ExampleNewReader() {
	r := NewReader("<html><head></head><body>hello</body></html>")
	doc, _ := html.Parse(r)
	html.Render(os.Stdout, doc)
	// Output: <html><head></head><body>hello</body></html>
}
