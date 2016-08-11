package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"sync"
)

func save(target *url.URL, r io.Reader) error {
	p := target.Path
	if p == "/" {
		p = "index.html"
	}

	f, err := os.Open(p)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}

	return nil
}

var wg sync.WaitGroup
var tokens = make(chan struct{}, 10)
var seen = make(map[string]bool)

func crawl(target *url.URL, depth int, list chan<- *url.URL, c canceler) {
	defer wg.Done()

	if depth == 0 || c.cancelled() {
		return
	}

	var buf bytes.Buffer
	tokens <- struct{}{} // acquire a token
	err := fetch(target, &buf, c)
	<-tokens
	if err != nil {
		return
	}

	links, err := extractLinks(target, &buf)
	if err != nil {
		return
	}

	for _, link := range links {
		if _, ok := seen[link.String()]; ok {
			continue
		}
		seen[link.String()] = true

		list <- link
		wg.Add(1)
		go crawl(link, depth-1, list, c)
	}
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

	list := make(chan *url.URL)
	c := newCanceller()
	wg.Add(1)
	go crawl(entry, *depth, list, c)

	// Canceller
	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("start cancel")
		c.cancel()
	}()

	// Closer
	go func() {
		wg.Wait()
		close(list)
	}()

	for link := range list {
		fmt.Println(link)
	}

	if c.cancelled() {
		fmt.Println("canceled")
	}
}
