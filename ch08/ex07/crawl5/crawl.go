package main

import (
	"flag"
	"fmt"
	"gopl.io/ch5/links"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

type work struct {
	list  []string
	depth int
}

var tokens = make(chan struct{}, 10)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	depth := flag.Int("depth", 3, "max depth of recursie crawling")
	flag.Parse()

	entry, err := url.Parse(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("depth=%d\n", *depth)
	fmt.Printf("Entry=%v\n", entry)
	fmt.Printf("Host=%v\n", entry.Host)

	worklist := make(chan work)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() {
		worklist <- work{[]string{entry.String()}, 0}
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		w := <-worklist
		if w.depth == *depth {
			continue
		}

		for _, link := range w.list {
			target, err := url.Parse(link)
			if err != nil {
				continue
			}

			if target.Host != entry.Host {
				continue
			}

			if seen[link] {
				continue
			}

			seen[link] = true
			n++
			go func(link string) {
				fetch(link)
				worklist <- work{crawl(link), w.depth + 1}
			}(link)
		}
	}
}
