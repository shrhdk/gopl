package main

import (
	"flag"
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

type work struct {
	list  []string
	depth int
}

var tokens = make(chan struct{}, 20)

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

	fmt.Printf("depth=%d\n", *depth)
	fmt.Printf("URL=%v\n", flag.Args())

	worklist := make(chan work)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() {
		worklist <- work{flag.Args(), 0}
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		w := <-worklist
		if w.depth == *depth {
			continue
		}

		for _, link := range w.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- work{crawl(link), w.depth + 1}
				}(link)
			}
		}
	}
}
