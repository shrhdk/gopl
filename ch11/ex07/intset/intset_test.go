package intset

import (
	"math/rand"
	"testing"
)

func randomSet(seed int64) map[int]struct{} {
	rng := rand.New(rand.NewSource(seed))
	set := make(map[int]struct{})
	for i := 0; i < 1000; i++ {
		set[rng.Intn(1000)] = struct{}{}
	}
	return set
}

func BenchmarkBitIntSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var set BitIntSet
		for i := range randomSet(int64(b.N)) {
			set.Add(i)
		}
	}
}

func BenchmarkBitIntSet32Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var set BitIntSet32
		for i := range randomSet(int64(b.N)) {
			set.Add(i)
		}
	}
}

func BenchmarkMapIntSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := NewMapIntSet()
		for i := range randomSet(int64(b.N)) {
			set.Add(i)
		}
	}
}

func BenchmarkBitIntSetUnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var set1, set2 BitIntSet
		for i := range randomSet(int64(b.N)) {
			set1.Add(i)
		}
		for i := range randomSet(int64(b.N)) {
			set2.Add(i)
		}
		set1.UnionWith(&set2)
	}
}

func BenchmarkBitInt32SetUnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var set1, set2 BitIntSet32
		for i := range randomSet(int64(b.N)) {
			set1.Add(i)
		}
		for i := range randomSet(int64(b.N)) {
			set2.Add(i)
		}
		set1.UnionWith(&set2)
	}
}

func BenchmarkMapIntSetUnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set1, set2 := NewMapIntSet(), NewMapIntSet()
		for i := range randomSet(int64(b.N)) {
			set1.Add(i)
		}
		for i := range randomSet(int64(b.N)) {
			set2.Add(i)
		}
		set1.UnionWith(set2)
	}
}
