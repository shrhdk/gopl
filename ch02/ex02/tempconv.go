package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string { return fmt.Sprintf("%.2f°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
