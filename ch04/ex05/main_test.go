package ex05

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	test(t, []string{}, []string{})

	test(t, []string{"A", "B", "C", "D"}, []string{"A", "B", "C", "D"})
	test(t, []string{"A", "B", "A", "C"}, []string{"A", "B", "A", "C"})
	test(t, []string{"A", "B", "C", "B"}, []string{"A", "B", "C", "B"})
	test(t, []string{"A", "B", "C", "A"}, []string{"A", "B", "C", "A"})

	test(t, []string{"A", "A", "B", "C"}, []string{"A", "B", "C"})
	test(t, []string{"A", "A", "A", "B"}, []string{"A", "B"})

	test(t, []string{"A", "B", "C", "C"}, []string{"A", "B", "C"})
	test(t, []string{"A", "B", "B", "B"}, []string{"A", "B"})

	test(t, []string{"A", "A", "A", "A"}, []string{"A"})
}

func test(t *testing.T, given, expected []string) {
	given = RemoveDuplication(given)
	if !reflect.DeepEqual(given, expected) {
		t.Errorf("\ngot  %v\nwant %v", given, expected)
	}
}
