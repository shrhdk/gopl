package ex03

import "testing"

func TestReverse(t *testing.T) {
	test(t, [4]int{0, 1, 2, 3}, [4]int{3, 2, 1, 0})
}

func test(t *testing.T, given, expected [4]int) {
	Reverse(&given)
	if given != expected {
		t.Errorf("\ngot  %c\nwant %c", given, expected)
	}
}
