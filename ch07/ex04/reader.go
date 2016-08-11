package main

import (
	"io"
)

type myReader struct {
	buf []byte
}

func (r *myReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.buf)
	r.buf = r.buf[n:]
	if len(r.buf) == 0 {
		return n, io.EOF
	}
	return n, nil
}

func NewReader(s string) io.Reader {
	return &myReader{[]byte(s)}
}
