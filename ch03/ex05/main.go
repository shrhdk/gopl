package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, xmax    = -2, 2
		ymin, ymax    = -2, 2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations uint16 = 1 << 9

	var v complex128
	for n := uint16(0); n < iterations; n++ {
		v = v*v + z

		r := uint8(((n >> 6) & 0x0007) / 7 * 255)
		g := uint8(((n >> 3) & 0x0007) / 7 * 255)
		b := uint8(((n >> 0) & 0x0007) / 7 * 255)

		if cmplx.Abs(v) > 2 {
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}
