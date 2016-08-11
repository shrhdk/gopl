package memo

import (
	"sync"
	"testing"
	"time"
)

func TestCancelGet(t *testing.T) {
	d := 500 * time.Millisecond

	m := New(func(key string) (interface{}, error) {
		time.Sleep(d)
		return key, nil
	})

	var d1, d2 time.Duration

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		start := time.Now()
		m.Get("dummy", NewCanceller())
		d1 = time.Since(start)
		wg.Done()
	}()

	c := NewCanceller()
	wg.Add(1)
	go func() {
		start := time.Now()
		m.Get("dummy", c)
		d2 = time.Since(start)
		wg.Done()
	}()

	c.Cancel()

	wg.Wait()

	if d1 < d {
		t.Errorf("first call will spent %v, but was %v", d, d1)
	}

	if d < d2 {
		t.Errorf("second invocation will not spent %v, but was %v", d, d2)
	}
}
