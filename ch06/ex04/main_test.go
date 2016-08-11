package intset

import (
	"reflect"
	"testing"
)

func TestElems(t *testing.T) {
	testElems(t)
	testElems(t, 0)
	testElems(t, 0, 63)
	testElems(t, 0, 63, 64)
	testElems(t, 0, 63, 64, 127)
}

func testElems(t *testing.T, given ...int) {
	var s IntSet
	s.AddAll(given...)
	actual := s.Elems()

	if !reflect.DeepEqual(actual, given) {
		t.Errorf("got %s\nwant %s", actual, given)
	}
}
