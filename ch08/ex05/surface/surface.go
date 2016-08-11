package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sync"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6 // 30deg
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	drawParallel(os.Stdout, 10)
	//draw(os.Stdout)
}

// Single

func draw(dst io.Writer) {
	fmt.Fprintf(dst, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: green; fill: black; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if p, err := polygon(i, j); err == nil {
				fmt.Fprintf(dst, p)
			}
		}
	}

	fmt.Fprintf(dst, "</svg>\n")
}

// Parallel

var wg sync.WaitGroup
var tokens chan struct{}

func drawParallel(dst io.Writer, num int) {
	fmt.Fprintf(dst, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: green; fill: black; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	tokens = make(chan struct{}, num)
	polygons := make([]string, cells)

	for i := 0; i < cells; i++ {
		wg.Add(1)
		go polygonOf(polygons, i)
	}

	wg.Wait()

	for _, p := range polygons {
		fmt.Fprintf(dst, p)
	}

	fmt.Fprintf(dst, "</svg>\n")
}

func polygonOf(polygons []string, i int) {
	defer wg.Done()

	tokens <- struct{}{}

	var buf bytes.Buffer
	for j := 0; j < cells; j++ {
		if p, err := polygon(i, j); err == nil {
			buf.WriteString(p)
		}
	}

	<-tokens

	polygons[i] = buf.String()
}

// Common

func polygon(i, j int) (string, error) {
	ax, ay, err := corner(i+1, j)
	if err != nil {
		return "", err
	}

	bx, by, err := corner(i, j)
	if err != nil {
		return "", err
	}

	cx, cy, err := corner(i, j+1)
	if err != nil {
		return "", err
	}

	dx, dy, err := corner(i+1, j+1)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy), nil
}

func corner(i, j int) (sx float64, sy float64, err error) {
	x := realize(i)
	y := realize(j)
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 1) || math.IsInf(z, -1) {
		err = errors.New("Invalid Result")
		return
	}

	sx, sy = project(x, y, z)
	return
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
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
