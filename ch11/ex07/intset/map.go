package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet struct {
	m map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{make(map[int]bool)}
}

func (s *MapIntSet) Add(x int) {
	s.m[x] = true
}

func (s *MapIntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for i := range t.m {
		s.Add(i)
	}
}

func (s *MapIntSet) IntersectWith(t *MapIntSet) {
	for i := range s.m {
		if !t.m[i] {
			delete(s.m, i)
		}
	}
}

func (s *MapIntSet) DifferenceWith(t *MapIntSet) {
	for i := range s.m {
		if t.m[i] {
			delete(s.m, i)
		}
	}
}

func (s *MapIntSet) SymmetricDifference(t *MapIntSet) {
	tmp := make(map[int]bool)
	for i := range s.m {
		if !t.m[i] {
			tmp[i] = true
		}
	}
	for i := range t.m {
		if !s.m[i] {
			tmp[i] = true
		}
	}
	s.m = tmp
}

func (s *MapIntSet) Elems() []int {
	var elems []int
	for i := range s.m {
		elems = append(elems, i)
	}
	sort.Sort(sort.IntSlice(elems))
	return elems
}

func (s *MapIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i := range s.Elems() {
		if buf.Len() > len("{") {
			buf.WriteRune(' ')
		}
		fmt.Fprint(&buf, i)
	}
	buf.WriteByte('}')
	return buf.String()
}
