package intset

import (
	"testing"
)

func TestIntersectWith(t *testing.T) {
	name := "IntersectWith"
	method := (*IntSet).IntersectWith

	test(t, name, method,
		[]int{},
		[]int{},
		[]int{})

	test(t, name, method,
		[]int{0},
		[]int{1},
		[]int{})

	test(t, name, method,
		[]int{0},
		[]int{64},
		[]int{})

	test(t, name, method,
		[]int{0},
		[]int{0},
		[]int{0})

	test(t, name, method,
		[]int{0, 63, 64},
		[]int{0, 63, 64},
		[]int{0, 63, 64})
}

func TestDifferenceWith(t *testing.T) {
	name := "DifferenceWith"
	method := (*IntSet).DifferenceWith

	test(t, name, method,
		[]int{},
		[]int{},
		[]int{})

	test(t, name, method,
		[]int{0},
		[]int{0},
		[]int{})

	test(t, name, method,
		[]int{0},
		[]int{64},
		[]int{0})

	test(t, name, method,
		[]int{0, 64},
		[]int{64},
		[]int{0})

	test(t, name, method,
		[]int{0, 63, 64},
		[]int{0, 64},
		[]int{63})
}

func TestSymmetricDifferenceWith(t *testing.T) {
	name := "SymmetricDifference"
	method := (*IntSet).SymmetricDifference

	test(t, name, method,
		[]int{},
		[]int{},
		[]int{})

	test(t, name, method,
		[]int{0},
		[]int{},
		[]int{0})

	test(t, name, method,
		[]int{0},
		[]int{63},
		[]int{0, 63})

	test(t, name, method,
		[]int{0},
		[]int{63, 64},
		[]int{0, 63, 64})

	test(t, name, method,
		[]int{0, 64},
		[]int{0, 63, 64},
		[]int{63})
}

func test(t *testing.T, name string, method func(s, t *IntSet), given1, given2, expected []int) {
	var s1, s2 IntSet
	s1.AddAll(given1...)
	s2.AddAll(given2...)

	method(&s1, &s2)

	var e IntSet
	e.AddAll(expected...)

	if s1.String() != e.String() {
		t.Errorf("\n%s %v %v\ngot %s\nwant %s", name, given1, given2, s1.String(), e.String())
	}
}
