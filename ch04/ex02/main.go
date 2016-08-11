package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	usage := fmt.Sprintf("Usage\n\t echo -n hello | %s [256|384|512]\n", os.Args[0])

	n := 256

	if len(os.Args) == 2 {
		n, _ = strconv.Atoi(os.Args[1])
	}

	bytes, _ := ioutil.ReadAll(os.Stdin)

	switch n {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256(bytes))
		os.Exit(0)
		return
	case 384:
		fmt.Printf("%x\n", sha512.Sum384(bytes))
		os.Exit(0)
		return
	case 512:
		fmt.Printf("%x\n", sha512.Sum512(bytes))
		os.Exit(0)
		return
	default:
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
		return
	}
}
