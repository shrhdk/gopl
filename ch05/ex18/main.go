package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	url := os.Args[1]

	local, n, err := fetch(url)
	if err != nil {
		log.Fatalf("failed to fetch %s", url)
	}

	fmt.Printf("fetch %s as %s (%dbytes)", url, local, n)
}

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
