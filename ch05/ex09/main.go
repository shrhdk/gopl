package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := expand("$foo $bar", func(t string) string {
		switch t {
		case "$foo":
			return "hello"
		case "$bar":
			return "world"
		default:
			return ""
		}
	})
	fmt.Println(s)
}

func expand(s string, f func(string) string) string {
	var b bytes.Buffer
	forEachWord(s, func(w string) {
		if strings.HasPrefix(w, "$") {
			b.WriteString(f(w))
		} else {
			b.WriteString(w)
		}
		b.WriteString(" ")
	})
	return b.String()
}

func forEachWord(s string, f func(string)) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		f(scanner.Text())
	}
}
