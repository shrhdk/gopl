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
		subSize                 = 2
		xmin, xmax              = -2, 2
		ymin, ymax              = -2, 2
		width, height           = 1024, 1024
		superWidth, superHeight = width * subSize, height * subSize
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for sy := 0; sy < superHeight; sy += subSize {
		for sx := 0; sx < superWidth; sx += subSize {
			var sum uint16

			for dy := 0; dy < subSize; dy++ {
				for dx := 0; dx < subSize; dx++ {
					x := float64(sx+dx)/superWidth*(xmax-xmin) + xmin
					y := float64(sy+dy)/superHeight*(ymax-ymin) + ymin
					z := complex(x, y)
					sum += uint16(mandelbrot(z))
				}
			}

			px := sx / subSize
			py := sy / subSize
			avg := uint8(sum / (subSize * subSize))

			img.Set(px, py, color.Gray{avg})
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) uint8 {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			return 255 - contrast*n
		}
	}
	return 0
}
