package links

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// filename
	filename := "." + resp.Request.URL.Path
	if strings.HasSuffix(filename, "/") {
		filename += "index.html"
	}

	// create directory
	dir := filepath.Dir(filename)
	_, err = os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, 0777)
	}

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to write response to file: %v", err)
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read wrote file: %v", err)
	}

	doc, err := html.Parse(f)
	f.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode {
			var link string

			switch n.Data {
			case "a":
				link = getAttr(n.Attr, "href")
			case "img":
				link = getAttr(n.Attr, "src")
			case "script":
				link = getAttr(n.Attr, "src")
			case "link":
				link = getAttr(n.Attr, "href")
			case "frame":
				link = getAttr(n.Attr, "src")
			case "iframe":
				link = getAttr(n.Attr, "src")
			}

			linkUrl, err := resp.Request.URL.Parse(link)
			if err != nil {
				return // ignore bad URLs
			}
			if linkUrl.Host != resp.Request.URL.Host {
				return // ignore external links
			}
			links = append(links, linkUrl.String())
		}
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

func getAttr(attrs []html.Attribute, name string) string {
	for _, a := range attrs {
		if a.Key == name {
			return a.Val
		}
	}
	return ""
}
