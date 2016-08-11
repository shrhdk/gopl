package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func fractal64(w io.Writer, psize int, rsize, cx, cy float32) {
	img := image.NewRGBA(image.Rect(0, 0, psize, psize))

	for py := 0; py < psize; py++ {
		y := realize32(py, psize, rsize) + cy
		for px := 0; px < psize; px++ {
			x := realize32(px, psize, rsize) + cx
			z := complex(x, y)

			count := mandelbrot64(z)
			color := color.Gray{255 - contrast*count}

			img.Set(px, py, color)
		}
	}

	png.Encode(w, img)
}

func realize32(d int, psize int, rsize float32) float32 {
	return float32(d)/float32(psize)*rsize - rsize/2
}

func mandelbrot64(z complex64) uint8 {
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > th {
			return n
		}
	}

	return 255
}
