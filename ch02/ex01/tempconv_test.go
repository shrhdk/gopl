// tempconv converts between C, F, and K.
package tempconv

import "testing"

func TestCToK(t *testing.T) {
	given := Celsius(0)
	expected := Kelvin(273.15)
	actual := CToK(given)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestKToC(t *testing.T) {
	given := Kelvin(273.15)
	expected := Celsius(0)
	actual := KToC(given)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
