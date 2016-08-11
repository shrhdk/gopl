package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	breathFirst(tree, os.Args[1:])
}

func tree(path string) []string {
	fmt.Println(path)

	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Print(err)
	}

	var dirs []string
	for _, fileInfo := range fileInfos {
		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		if fileInfo.IsDir() {
			dirs = append(dirs, filepath.Join(path, fileInfo.Name()))
		}
	}

	return dirs
}

func breathFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
