package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popCount(word)
	}
	return count
}

func (s *IntSet) Add(x int) {
	word, bit := pos(x)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Remove(x int) {
	word, bit := pos(x)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
	s.shrink()
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() IntSet {
	var c IntSet
	c.words = make([]uint64, len(s.words))
	copy(c.words, s.words)
	return c
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func pos(x int) (word int, bit uint) {
	return x / 64, uint(x % 64)
}

func (s *IntSet) shrink() {
	var i int
	for i = len(s.words) - 1; i >= 0; i-- {
		if s.words[i] != 0 {
			break
		}
	}
	s.words = s.words[:i+1]
}

func popCount(x uint64) (count int) {
	for x != 0 {
		x &= (x - 1)
		count++
	}
	return count
}
