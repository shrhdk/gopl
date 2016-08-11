package ex15

import (
	"testing"
)

func TestMin(t *testing.T) {
	test(t, 1, min(1))
	test(t, 1, min(1, 2, 3))
	test(t, 4, min(6, 5, 4))
	test(t, 3, min(9, 3, 6))
	test(t, 1, min(1, 1, 1))
}

func TestMax(t *testing.T) {
	test(t, 1, max(1))
	test(t, 3, max(1, 2, 3))
	test(t, 6, max(6, 5, 4))
	test(t, 9, max(6, 9, 3))
	test(t, 1, max(1, 1, 1))
}

func TestMin2(t *testing.T) {
	test(t, 1, min2(1))
	test(t, 1, min2(1, 2, 3))
	test(t, 4, min2(6, 5, 4))
	test(t, 3, min2(9, 3, 6))
	test(t, 1, min2(1, 1, 1))
}

func TestMax2(t *testing.T) {
	test(t, 1, max2(1))
	test(t, 3, max2(1, 2, 3))
	test(t, 6, max2(6, 5, 4))
	test(t, 9, max2(6, 9, 3))
	test(t, 1, max2(1, 1, 1))
}

func test(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
