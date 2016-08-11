package ex07

import (
	"golang.org/x/net/html"
	"os"
	"testing"
)

func TestElementsByTagName(t *testing.T) {
	f, err := os.Open("test.html")
	if err != nil {
		t.Errorf("failed to open test file: %v", err)
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		t.Errorf("failed to parse test file: %v", err)
	}

	elems := ElementsByTagName(doc, "body", "li")

	for _, elem := range elems {
		t.Log(elem.Data)
	}
}
