package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

// go build github.com/shrhdk/gopl/ch10/ex01/imgconv
// cat gopher.jpg | ./imgconv -t png > gopher.png
func main() {
	format := flag.String("t", "jpeg", "[jpg|jpeg|png|gif]")
	flag.Parse()

	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	switch *format {
	case "jpg", "jpeg":
		err = jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: 95})
	case "png":
		err = png.Encode(os.Stdout, img)
	case "gif":
		err = gif.Encode(os.Stdout, img, nil)
	default:
		fmt.Fprintf(os.Stderr, "Unknown format: %s\n", *format)
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}
}
