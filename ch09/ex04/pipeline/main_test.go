package pipeline

import (
	"fmt"
	"testing"
	"time"
)

func TestPipeline1000(t *testing.T) {
	for _, size := range []int{10, 100, 1000, 10000, 50000, 100000, 1000000} {
		in, out := newPipeline(size)

		start := time.Now()
		in <- "hello"
		<-out
		fmt.Printf("%d\t%d\n", size, time.Since(start))
	}
}
