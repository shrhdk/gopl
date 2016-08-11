package popcount

import (
	"testing"
)

// Test

func TestPopCount(t *testing.T) {
	var given uint64
	var expected, actual int

	given = 0
	expected = 0
	actual = PopCount4(given)
	if actual != expected {
		t.Errorf("got %v\nwant %v (input = %d)", actual, expected, given)
	}

	given = 1
	expected = 1
	actual = PopCount4(given)
	if actual != expected {
		t.Errorf("got %v\nwant %v (input = %d)", actual, expected, given)
	}

	given = 255
	expected = 8
	actual = PopCount4(given)
	if actual != expected {
		t.Errorf("got %v\nwant %v (input = %d)", actual, expected, given)
	}
}

// Benchmark

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

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(1)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount4(1)
	}
}
