package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	host = "localhost"
	port = 8000
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	fmt.Printf("Listening at %s:%d...", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	x := parseQuery(r.URL.Query(), "x", 0.0)
	y := parseQuery(r.URL.Query(), "y", 0.0)
	scale := parseQuery(r.URL.Query(), "scale", 1.0)

	fractal(w, x, y, scale)
}

func parseQuery(values url.Values, key string, fallback float64) float64 {
	if values[key] == nil || len(values[key]) < 0 {
		return fallback
	}

	v, err := strconv.ParseFloat(values[key][0], 64)

	if err != nil {
		return fallback
	}

	return v
}

const (
	contrast = 15
	th       = 1
	size     = 1024
)

func fractal(w io.Writer, cx float64, cy float64, scale float64) {
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for py := 0; py < size; py++ {
		y := realize(py, scale) + cy
		for px := 0; px < size; px++ {
			x := realize(px, scale) + cx
			z := complex(x, y)

			count := solve(z)
			color := color.Gray{uint8(count) * contrast}

			img.Set(px, py, color)
		}
	}

	png.Encode(w, img)
}

func realize(d int, scale float64) float64 {
	width := 10 / scale
	return float64(d)/size*width - width/2
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
