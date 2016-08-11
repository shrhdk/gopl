package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

const bitLength = 32 << (^uint(0) >> 63)

func (s *IntSet) Add(x int) {
	word, bit := pos(x)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitLength; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitLength*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func pos(x int) (word int, bit uint) {
	return x / bitLength, uint(x % bitLength)
}
