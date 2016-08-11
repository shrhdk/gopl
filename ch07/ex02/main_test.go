package ex02

import (
	"bytes"
	"testing"
)

func TestCountingWrite(t *testing.T) {
	tests := []struct {
		s string
		c int64
	}{
		{"", 0},
		{"12345", 5},
		{"123456789", 9},
	}

	for _, test := range tests {
		var b bytes.Buffer
		w, c := CountingWriter(&b)

		// act
		w.Write([]byte(test.s))

		// verify
		if b.String() != test.s {
			t.Errorf("expected string is %s, but actual is %s", b.String(), test.s)
		}

		if *c != test.c {
			t.Errorf("expected count is %d, but actual is %d", test.c, *c)
		}
	}
}
