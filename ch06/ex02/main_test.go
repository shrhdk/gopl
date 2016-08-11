package intset

import "testing"

func TestAddAll(t *testing.T) {
	testAddAll(t, "{}")
	testAddAll(t, "{0}", 0)
	testAddAll(t, "{0 63}", 0, 63)
	testAddAll(t, "{0 63 64}", 0, 63, 64)
	testAddAll(t, "{0 63 64 127}", 0, 63, 64, 127)
}

func testAddAll(t *testing.T, expected string, given ...int) {
	var actual IntSet
	actual.AddAll(given...)

	if actual.String() != expected {
		t.Errorf("got %s\nwant %s", actual.String(), expected)
	}
}
