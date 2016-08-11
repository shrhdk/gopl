package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.1
	angle         = math.Pi / 6 // 30deg
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var max, min float64 = -math.MaxFloat64, math.MaxFloat64

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%d' height='%d'>\n", width, height)

	calcMinMax()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			c := surfaceColor(float64(i)+0.5, float64(j)+0.5)

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill: %s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}

	fmt.Println("</svg>")
}

func calcMinMax() {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange*(float64(i)+0.5)/cells - 0.5
			y := xyrange*(float64(j)+0.5)/cells - 0.5

			z := f(x, y)

			if z < min {
				min = z
			}

			if z > max {
				max = z
			}
		}
	}
}

func surfaceColor(x, y float64) string {
	x = xyrange * (x/cells - 0.5)
	y = xyrange * (y/cells - 0.5)

	z := (f(x, y) - min) / (max - min)

	r := uint8(z * 255)
	g := 0x00
	b := uint8((1 - z) * 255)

	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func corner(i, j int) (float64, float64) {
	x := realize(i)
	y := realize(j)
	z := f(x, y)

	return project(x, y, z)
}

func realize(i int) float64 {
	return xyrange * (float64(i)/cells - 0.5)
}

func project(x, y, z float64) (sx, sy float64) {
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	x = math.Cos(x / 2)
	y = math.Cos(y / 2)

	return 0.5 * (x + y)
}
