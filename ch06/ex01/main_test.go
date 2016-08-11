package intset

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	assertLenEquals(t, intset(), 0)
	assertLenEquals(t, intset(0), 1)
	assertLenEquals(t, intset(0, 63), 2)
	assertLenEquals(t, intset(0, 63, 64), 3)
	assertLenEquals(t, intset(0, 63, 64, 127), 4)
	assertLenEquals(t, intset(0, 63, 64, 127, 128), 5)
	assertLenEquals(t, intset(0, 0), 1)
}

func TestRemoveNothing(t *testing.T) {
	var s IntSet
	s.Remove(1)
	assertIntSetEquals(t, s, intset())
}

func TestRemove1(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Remove(1)
	assertIntSetEquals(t, s, intset())
}

func TestRemove2(t *testing.T) {
	var s IntSet
	s.Add(2)
	s.Add(64)
	s.Remove(2)
	assertIntSetEquals(t, s, intset(64))
}

func TestRemove64(t *testing.T) {
	var s IntSet
	s.Add(2)
	s.Add(64)
	s.Remove(64)
	assertIntSetEquals(t, s, intset(2))
}

func TestClearNothing(t *testing.T) {
	var s IntSet
	s.Clear()
	assertIntSetEquals(t, s, intset())
}

func TestClear1(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Clear()
	assertIntSetEquals(t, s, intset())
}

func TestClear2(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(2)
	s.Clear()
	assertIntSetEquals(t, s, intset())
}

func TestCopyEmpty(t *testing.T) {
	var s IntSet
	c := s.Copy()
	s.Add(1)
	assertIntSetEquals(t, c, intset())
}

func TestCopy1(t *testing.T) {
	var s IntSet
	s.Add(1)
	c := s.Copy()
	s.Remove(1)
	assertIntSetEquals(t, c, intset(1))
}

func intset(xs ...int) IntSet {
	var s IntSet
	for _, x := range xs {
		s.Add(x)
	}
	return s
}

func assertLenEquals(t *testing.T, actual IntSet, expected int) {
	if actual.Len() != expected {
		t.Errorf("got %s\nwant %s", actual.Len(), expected)
	}
}

func assertIntSetEquals(t *testing.T, actual, expected IntSet) {
	if len(actual.words) == 0 && len(expected.words) == 0 {
		return
	}

	if !reflect.DeepEqual(actual.words, expected.words) {
		t.Errorf("got %s\nwant %s", actual.String(), expected.String())
	}
}
