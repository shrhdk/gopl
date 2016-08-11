package intset

type IntSet struct {
	words []uint64
}

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

func (s *IntSet) Elems() []int {
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

func pos(x int) (word int, bit uint) {
	return x / 64, uint(x % 64)
}
