package main

import (
	"bytes"
	"testing"
)

func BenchmarkDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		draw(&buf)
	}
}

func BenchmarkDrawParallel2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		drawParallel(&buf, 2)
	}
}

func BenchmarkDrawParallel3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		drawParallel(&buf, 3)
	}
}

func BenchmarkDrawParallel4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		drawParallel(&buf, 4)
	}
}

func BenchmarkDrawParallel5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		drawParallel(&buf, 5)
	}
}

func BenchmarkDrawParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		drawParallel(&buf, 10)
	}
}

func BenchmarkDrawParallel50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		drawParallel(&buf, 50)
	}
}

func BenchmarkDrawParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		drawParallel(&buf, 100)
	}
}
