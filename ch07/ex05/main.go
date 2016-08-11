package ex05

import (
	"fmt"
	"io"
)

type limitReader struct {
	base  io.Reader
	limit int64
	count int64
}

func (lr *limitReader) remain() int64 {
	return lr.limit - lr.count
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	if int64(len(p)) >= lr.remain() {
		p = p[:lr.remain()]
	}

	n, err = lr.base.Read(p)
	lr.count += int64(n)

	if err == nil && lr.remain() <= 0 {
		err = io.EOF
	}

	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	if r == nil {
		panic("r is nil.")
	}

	if n < 0 {
		panic(fmt.Sprintf("n must be 0 or over, but given n is %v", n))
	}

	return &limitReader{r, n, 0}
}
