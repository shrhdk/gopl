package ex02

import "io"

type countingWriter struct {
	writer io.Writer
	count  int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p)
	cw.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw countingWriter
	cw.writer = w
	return &cw, &(cw.count)
}
