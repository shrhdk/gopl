package intset

import (
	"bytes"
	"fmt"
)

type BitIntSet struct {
	words []uint64
}

func (s *BitIntSet) Add(x int) {
	word, bit := pos(x)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *BitIntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *BitIntSet) UnionWith(t *BitIntSet) {
	s.operate(t, func(x, y uint64) uint64 {
		return x | y
	})
}

func (s *BitIntSet) IntersectWith(t *BitIntSet) {
	s.operate(t, func(x, y uint64) uint64 {
		return x & y
	})
}

func (s *BitIntSet) DifferenceWith(t *BitIntSet) {
	s.operate(t, func(x, y uint64) uint64 {
		return x &^ y
	})
}

func (s *BitIntSet) SymmetricDifference(t *BitIntSet) {
	s.operate(t, func(x, y uint64) uint64 {
		return x ^ y
	})
}

func (s *BitIntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return elems
}

func (s *BitIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteRune(' ')
				}
				fmt.Fprint(&buf, 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *BitIntSet) operate(t *BitIntSet, f func(s, t uint64) uint64) {
	l := max(len(s.words), len(t.words))
	var ws []uint64

	for i := 0; i < l; i++ {
		ws = append(ws, f(s.getWord(i), t.getWord(i)))
	}

	s.words = ws
	s.shrink()
}

func (s *BitIntSet) getWord(i int) uint64 {
	if len(s.words) <= i {
		return 0
	}

	return s.words[i]
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func pos(x int) (word int, bit uint) {
	return x / 64, uint(x % 64)
}

func (s *BitIntSet) shrink() {
	var i int
	for i = len(s.words) - 1; i >= 0; i-- {
		if s.words[i] != 0 {
			break
		}
	}
	s.words = s.words[:i+1]
}
