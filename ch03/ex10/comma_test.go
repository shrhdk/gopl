package ex10

import "testing"

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

func test(t *testing.T, given, expected string) {
	actual := Comma(given)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
