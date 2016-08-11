package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputs []string = os.Args[1:]

	if len(inputs) == 0 {
		var value string
		fmt.Scan(&value)
		inputs = strings.Split(value, " ")
	}

	for _, arg := range inputs {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		printconv(v)
	}
}

func printconv(v float64) {
	f := Fahrenheit(v)
	c := Celsius(v)
	fmt.Printf("%s = %s, %s = %s\n", c, CToF(c), f, FToC(f))

	kg := Kilogram(v)
	p := Pound(v)
	fmt.Printf("%s = %s, %s = %s\n", kg, KilogramToPound(kg), p, PoundToKilogram(p))

	m := Meter(v)
	ft := Feet(v)
	fmt.Printf("%s = %s, %s = %s\n", m, MeterToFeet(m), ft, FeetToMeter(ft))
}
