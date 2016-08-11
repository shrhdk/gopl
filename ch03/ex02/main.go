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

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: green; fill: black; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := realize(i)
	y := realize(j)
	z := egg(x, y)

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

func egg(x, y float64) float64 {
	x = math.Cos(x)
	y = math.Cos(y)

	if x <= 0 || y <= 0 {
		x, y = 0, 0
	}

	return 0.3 * (x + y)
}
