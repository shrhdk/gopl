package popcount

import (
	"fmt"
	"testing"
	"time"
)

func TestPopCount(t *testing.T) {
	tests := []struct {
		X uint64
		C int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{255, 8},
	}

	for _, test := range tests {
		act := PopCount(test.X)
		if act != test.C {
			t.Errorf("want %d for %d, got %d", test.C, test.X, act)
		}
	}
}

// v option of `go test` shows result of
// go test -v github.com/shrhdk/gopl/ch09/ex02/popcount
func TestPopCountPerformance(t *testing.T) {
	for i := 0; i < 10; i++ {
		start := time.Now()
		PopCount(0)
		du := time.Since(start)
		fmt.Printf("%d: %d\n", i, du)
	}
}
