package intset

import (
	"bytes"
	"fmt"
)

type BitIntSet32 struct {
	words []uint32
}

func (s *BitIntSet32) Add(x int) {
	word, bit := pos32(x)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *BitIntSet32) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *BitIntSet32) UnionWith(t *BitIntSet32) {
	s.operate(t, func(x, y uint32) uint32 {
		return x | y
	})
}

func (s *BitIntSet32) IntersectWith(t *BitIntSet32) {
	s.operate(t, func(x, y uint32) uint32 {
		return x & y
	})
}

func (s *BitIntSet32) DifferenceWith(t *BitIntSet32) {
	s.operate(t, func(x, y uint32) uint32 {
		return x &^ y
	})
}

func (s *BitIntSet32) SymmetricDifference(t *BitIntSet32) {
	s.operate(t, func(x, y uint32) uint32 {
		return x ^ y
	})
}

func (s *BitIntSet32) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 32*i+j)
			}
		}
	}
	return elems
}

func (s *BitIntSet32) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteRune(' ')
				}
				fmt.Fprint(&buf, 32*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *BitIntSet32) operate(t *BitIntSet32, f func(s, t uint32) uint32) {
	l := max32(len(s.words), len(t.words))
	var ws []uint32

	for i := 0; i < l; i++ {
		ws = append(ws, f(s.getWord(i), t.getWord(i)))
	}

	s.words = ws
	s.shrink()
}

func (s *BitIntSet32) getWord(i int) uint32 {
	if len(s.words) <= i {
		return 0
	}

	return s.words[i]
}

func max32(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func pos32(x int) (word int, bit uint) {
	return x / 32, uint(x % 32)
}

func (s *BitIntSet32) shrink() {
	var i int
	for i = len(s.words) - 1; i >= 0; i-- {
		if s.words[i] != 0 {
			break
		}
	}
	s.words = s.words[:i+1]
}
