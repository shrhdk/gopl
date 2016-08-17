package intset

import (
	"reflect"
	"testing"
)

var tests = []struct {
	s1 []int
	s2 []int
}{
	{[]int{}, []int{}},
	{[]int{0}, []int{1}},
	{[]int{63}, []int{0}},
	{[]int{63}, []int{0}},
	{[]int{0}, []int{63}},
	{[]int{0}, []int{64}},
	{[]int{0}, []int{0}},
	{[]int{0, 63, 64}, []int{0, 63, 64}},
	{[]int{0, 63, 64}, []int{0, 64}},
	{[]int{0, 64}, []int{0, 63, 64}},
}

func TestIntersectWith(t *testing.T) {
	for _, test := range tests {
		var bis1, bis2 BitIntSet
		bis1.AddAll(test.s1...)
		bis2.AddAll(test.s2...)
		bis1.IntersectWith(&bis2)

		var mis1, mis2 = NewMapIntSet(), NewMapIntSet()
		mis1.AddAll(test.s1...)
		mis2.AddAll(test.s2...)
		mis1.IntersectWith(mis2)

		if !reflect.DeepEqual(bis1.Elems(), mis1.Elems()) {
			t.Errorf("BitIntSet return %v, MapIntSet return %v for Intersect of %v and %v", bis1, mis1, test.s1, test.s2)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	for _, test := range tests {
		var bis1, bis2 BitIntSet
		bis1.AddAll(test.s1...)
		bis2.AddAll(test.s2...)
		bis1.DifferenceWith(&bis2)

		var mis1, mis2 = NewMapIntSet(), NewMapIntSet()
		mis1.AddAll(test.s1...)
		mis2.AddAll(test.s2...)
		mis1.DifferenceWith(mis2)

		if !reflect.DeepEqual(bis1.Elems(), mis1.Elems()) {
			t.Errorf("BitIntSet return %v, MapIntSet return %v for Difference of %v and %v", bis1, mis1, test.s1, test.s2)
		}
	}
}

func TestSymmetricDifferenceWith(t *testing.T) {
	for _, test := range tests {
		var bis1, bis2 BitIntSet
		bis1.AddAll(test.s1...)
		bis2.AddAll(test.s2...)
		bis1.SymmetricDifference(&bis2)

		var mis1, mis2 = NewMapIntSet(), NewMapIntSet()
		mis1.AddAll(test.s1...)
		mis2.AddAll(test.s2...)
		mis1.SymmetricDifference(mis2)

		if !reflect.DeepEqual(bis1.Elems(), mis1.Elems()) {
			t.Errorf("BitIntSet return %v, MapIntSet return %v for SymmetricDifference of %v and %v", bis1, mis1, test.s1, test.s2)
		}
	}
}
