package main

import (
	"fmt"
	"io"
	"os"

	"github.com/shrhdk/gopl/ch10/ex02/archive"
	_ "github.com/shrhdk/gopl/ch10/ex02/archive/tar"
	_ "github.com/shrhdk/gopl/ch10/ex02/archive/zip"
)

func main() {
	name := os.Args[1]

	f, err := os.Open(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	stat, err := f.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	r := io.NewSectionReader(f, 0, stat.Size())
	ar, err := archive.NewReader(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	visitAll(ar, func(fileInfo os.FileInfo, r io.Reader) {
		fmt.Println(fileInfo.Name())
	})
}

func visitAll(r archive.Reader, f func(f os.FileInfo, r io.Reader)) error {
	info, err := r.Next()
	if err == io.EOF {
		return nil
	} else if err != nil {
		return err
	}

	f(info, r)

	return visitAll(r, f)
}
