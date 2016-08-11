package ex01

import (
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		str string
		w   int
		l   int
	}{
		{"", 0, 1},
		{"a b c\n\ndef\ng", 5, 4},
	}

	for _, test := range tests {
		var c WordLineCounter
		c.Write([]byte(test.str))

		if c.Words != test.w {
			t.Errorf("expected words is %d, but actual is %d", test.w, c.Words)
		}

		if c.Lines != test.l {
			t.Errorf("expected lines is %d, but actual is %d", test.l, c.Lines)
		}
	}
}
