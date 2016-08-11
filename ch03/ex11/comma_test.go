package ex11

import "testing"

// Simple Integer

func TestComma1(t *testing.T) {
	test(t, "1", "1")
}

func TestComma2(t *testing.T) {
	test(t, "12", "12")
}

func TestComma3(t *testing.T) {
	test(t, "123", "123")
}

func TestComma4(t *testing.T) {
	test(t, "1234", "1,234")
}

func TestComma5(t *testing.T) {
	test(t, "12345", "12,345")
}

func TestComma6(t *testing.T) {
	test(t, "123456", "123,456")
}

func TestComma7(t *testing.T) {
	test(t, "1234567", "1,234,567")
}

// Negative Integer

func TestNegativeComma1(t *testing.T) {
	test(t, "-1", "-1")
}

func TestNegativeComma2(t *testing.T) {
	test(t, "-12", "-12")
}

func TestNegativeComma3(t *testing.T) {
	test(t, "-123", "-123")
}

func TestNegativeComma4(t *testing.T) {
	test(t, "-1234", "-1,234")
}

func TestNegativeComma5(t *testing.T) {
	test(t, "-12345", "-12,345")
}

func TestNegativeComma6(t *testing.T) {
	test(t, "-123456", "-123,456")
}

func TestNegativeComma7(t *testing.T) {
	test(t, "-1234567", "-1,234,567")
}

// Positive Integer

func TestPositiveComma1(t *testing.T) {
	test(t, "+1", "+1")
}

func TestPositiveComma2(t *testing.T) {
	test(t, "+12", "+12")
}

func TestPositiveComma3(t *testing.T) {
	test(t, "+123", "+123")
}

func TestPositiveComma4(t *testing.T) {
	test(t, "+1234", "+1,234")
}

func TestPositiveComma5(t *testing.T) {
	test(t, "+12345", "+12,345")
}

func TestPositiveComma6(t *testing.T) {
	test(t, "+123456", "+123,456")
}

func TestPositiveComma7(t *testing.T) {
	test(t, "+1234567", "+1,234,567")
}

// Simple Float

func TestFloatComma1(t *testing.T) {
	test(t, "1.000", "1.000")
}

func TestFloatComma2(t *testing.T) {
	test(t, "12.000", "12.000")
}

func TestFloatComma3(t *testing.T) {
	test(t, "123.000", "123.000")
}

func TestFloatComma4(t *testing.T) {
	test(t, "1234.000", "1,234.000")
}

func TestFloatComma5(t *testing.T) {
	test(t, "12345.000", "12,345.000")
}

func TestFloatComma6(t *testing.T) {
	test(t, "123456.000", "123,456.000")
}

func TestFloatComma7(t *testing.T) {
	test(t, "1234567.000", "1,234,567.000")
}

// Negative Float

func TestNegativeFloatComma1(t *testing.T) {
	test(t, "-1.000", "-1.000")
}

func TestNegativeFloatComma2(t *testing.T) {
	test(t, "-12.000", "-12.000")
}

func TestNegativeFloatComma3(t *testing.T) {
	test(t, "-123.000", "-123.000")
}

func TestNegativeFloatComma4(t *testing.T) {
	test(t, "-1234.000", "-1,234.000")
}

func TestNegativeFloatComma5(t *testing.T) {
	test(t, "-12345.000", "-12,345.000")
}

func TestNegativeFloatComma6(t *testing.T) {
	test(t, "-123456.000", "-123,456.000")
}

func TestNegativeFloatComma7(t *testing.T) {
	test(t, "-1234567.000", "-1,234,567.000")
}

// Positive Float

func TestPositiveFloatComma1(t *testing.T) {
	test(t, "+1.000", "+1.000")
}

func TestPositiveFloatComma2(t *testing.T) {
	test(t, "+12.000", "+12.000")
}

func TestPositiveFloatComma3(t *testing.T) {
	test(t, "+123.000", "+123.000")
}

func TestPositiveFloatComma4(t *testing.T) {
	test(t, "+1234.000", "+1,234.000")
}

func TestPositiveFloatComma5(t *testing.T) {
	test(t, "+12345.000", "+12,345.000")
}

func TestPositiveFloatComma6(t *testing.T) {
	test(t, "+123456.000", "+123,456.000")
}

func TestPositiveFloatComma7(t *testing.T) {
	test(t, "+1234567.000", "+1,234,567.000")
}

// Helper

func test(t *testing.T, given, expected string) {
	actual := Comma(given)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
