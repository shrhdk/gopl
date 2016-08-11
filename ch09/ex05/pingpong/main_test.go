package pingpong

import (
	"fmt"
	"testing"
	"time"
)

func TestPingPong(t *testing.T) {
	done := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()
	c := pingpong(done)
	fmt.Printf("pingpong: %d\n", c)
}
