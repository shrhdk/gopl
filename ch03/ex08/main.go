package main

import (
	"fmt"
	"os"
)

const (
	contrast   = 15
	iterations = 200
	th         = 2
)

// TODO: add big.Float version
// TODO: add big.Rat version

func main() {
	const (
		psize = 1024
		rsize = 1e-12
		cx    = -1.5
		cy    = 0
	)

	var out *os.File

	fmt.Println("Writing to complex64.png ...")
	out, _ = os.Create("./complex64.png")
	fractal64(out, psize, rsize, cx, cy)
	out.Close()

	fmt.Println("Writing to complex128.png ...")
	out, _ = os.Create("./complex128.png")
	fractal128(out, psize, rsize, cx, cy)
	out.Close()
}
