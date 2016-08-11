package popcount

import (
	"testing"
)

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(1)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(1)
	}
}
