package main

import (
	"fmt"
)

type Kilogram float64
type Pound float64

func KilogramToPound(kg Kilogram) Pound {
	return Pound(kg * 2.2046)
}

func PoundToKilogram(p Pound) Kilogram {
	return Kilogram(p / 2.2046)
}

func (kg Kilogram) String() string { return fmt.Sprintf("%.2fkg", kg) }

func (p Pound) String() string { return fmt.Sprintf("%.2flb", p) }
