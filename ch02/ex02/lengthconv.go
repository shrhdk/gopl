package main

import (
	"fmt"
)

type Meter float64
type Feet float64

func MeterToFeet(m Meter) Feet {
	return Feet(m * 3.2808)
}

func FeetToMeter(ft Feet) Meter {
	return Meter(ft / 3.2808)
}

func (m Meter) String() string { return fmt.Sprintf("%.2fm", m) }

func (ft Feet) String() string { return fmt.Sprintf("%.2fft", ft) }
