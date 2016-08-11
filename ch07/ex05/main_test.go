package ex05

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

type mock struct {
	size, pos int
	nErr      int
	err       error
}

func (m *mock) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		if m.nErr != -1 && m.pos == m.nErr {
			return i, m.err
		}

		if m.pos == m.size {
			return i, io.EOF
		}

		p[i] = byte(m.pos)
		m.pos++
	}

	return len(p), nil
}

// r == nil
func TestNilReader(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("no panic with nil reader.")
		}
	}()

	LimitReader(nil, 0)
}

// n < 0
func TestNegativeN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("no panic with negative n.")
		}
	}()

	LimitReader(&mock{0, 0, -1, nil}, -1)
}

var err = fmt.Errorf("error returned by mock")

type action struct {
	lenP        int
	expectedN   int
	expectedErr error
	expectedP   []byte
}

var tests = []struct {
	reader io.Reader
	acts   []action
}{
	{
		// len(r) < n
		LimitReader(&mock{4, 0, -1, nil}, 5),
		[]action{
			{5, 4, io.EOF, []byte{0, 1, 2, 3}},
		},
	},
	{
		// len(r) == n
		LimitReader(&mock{5, 0, -1, nil}, 5),
		[]action{
			{5, 5, io.EOF, []byte{0, 1, 2, 3, 4}},
		},
	},
	{
		// len(r) > n
		LimitReader(&mock{6, 0, -1, nil}, 5),
		[]action{
			{6, 5, io.EOF, []byte{0, 1, 2, 3, 4}},
		},
	},
	{
		// n == 0
		LimitReader(&mock{0, 0, -1, nil}, 5),
		[]action{
			{5, 0, io.EOF, []byte{}},
		},
	},
	{
		// Read 2 times
		LimitReader(&mock{6, 0, -1, nil}, 5),
		[]action{
			{3, 3, nil, []byte{0, 1, 2}},
			{3, 2, io.EOF, []byte{3, 4}},
		},
	},
	{
		// Base io.Reader returns error
		LimitReader(&mock{5, 0, 1, err}, 5),
		[]action{
			{5, 1, err, []byte{0}},
		},
	},
}

func TestLimitReader(t *testing.T) {
	for _, test := range tests {
		for _, act := range test.acts {
			// Act
			p := make([]byte, act.lenP)
			n, err := test.reader.Read(p)

			// Verify
			if n != act.expectedN {
				t.Errorf("n = %v, want %v", n, act.expectedN)
			}
			if err != act.expectedErr {
				t.Errorf("err = %v, want %v", err, act.expectedErr)
			}
			if !reflect.DeepEqual(act.expectedP, p[:n]) {
				t.Errorf("err = %v, want %v", p[:n], act.expectedP)
			}
		}
	}
}
