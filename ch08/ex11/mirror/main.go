package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	c := newCanceller()

	for _, target := range os.Args[1:] {
		wg.Add(1)
		go save(target, c)
	}

	wg.Wait()
}

func save(target string, c canceler) {
	defer wg.Done()

	u, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Replace(u.Host+u.Path, "/", "_", -1)

	var buf bytes.Buffer

	fetch(u, &buf, c)

	if c.cancelled() {
		fmt.Printf("cancel %s\n", target)
		return
	}

	fmt.Printf("complete %s\n", target)

	c.cancel()

	f, err := os.Create(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.Copy(f, &buf)
}
