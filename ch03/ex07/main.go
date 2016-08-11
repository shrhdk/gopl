package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	contrast      = 15
	th            = 1
	xmin, xmax    = -10, 10
	ymin, ymax    = -10, 10
	width, height = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			count := solve(z)
			color := color.Gray{uint8(count) * contrast}

			img.Set(px, py, color)
		}
	}

	png.Encode(os.Stdout, img)
}

func solve(zn complex128) (ct uint64) {
	for cmplx.Abs(zn) > th {
		zn = zn - f(zn)/df(zn)
		ct++
	}

	return
}

func f(z complex128) complex128 {
	return cmplx.Pow(z, 4) - 1
}

func df(z complex128) complex128 {
	return 4 * cmplx.Pow(z, 3)
}
