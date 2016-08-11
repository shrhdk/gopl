package ex04

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	test(t, []int{0, 1, 2, 3}, 0, []int{0, 1, 2, 3})
	test(t, []int{0, 1, 2, 3}, 1, []int{3, 0, 1, 2})
	test(t, []int{0, 1, 2, 3}, 2, []int{2, 3, 0, 1})
	test(t, []int{0, 1, 2, 3}, 3, []int{1, 2, 3, 0})
	test(t, []int{0, 1, 2, 3}, 4, []int{0, 1, 2, 3})
}

func test(t *testing.T, given []int, n int, expected []int) {
	Rotate(given, n)
	if !reflect.DeepEqual(given, expected) {
		t.Errorf("\ngot  %v\nwant %v", given, expected)
	}
}
