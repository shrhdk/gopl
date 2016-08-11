package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := wordfreq(os.Stdin)

	fmt.Printf("word\tcount\n")
	for s, n := range counts {
		fmt.Printf("%s\t%d\n", s, n)
	}
}

func wordfreq(r io.Reader) map[string]int {
	in := bufio.NewScanner(r)
	in.Split(bufio.ScanWords)

	counts := make(map[string]int)
	for in.Scan() {
		counts[in.Text()]++
	}

	return counts
}
