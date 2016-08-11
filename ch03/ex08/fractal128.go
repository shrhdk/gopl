package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func fractal128(w io.Writer, psize int, rsize, cx, cy float64) {
	img := image.NewRGBA(image.Rect(0, 0, psize, psize))

	for py := 0; py < psize; py++ {
		y := realize64(py, psize, rsize) + cy
		for px := 0; px < psize; px++ {
			x := realize64(px, psize, rsize) + cx
			z := complex(x, y)

			count := mandelbrot128(z)
			color := color.Gray{255 - contrast*count}

			img.Set(px, py, color)
		}
	}

	png.Encode(w, img)
}

func realize64(d int, psize int, rsize float64) float64 {
	return float64(d)/float64(psize)*rsize - rsize/2
}

func mandelbrot128(z complex128) uint8 {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > th {
			return n
		}
	}

	return 255
}
