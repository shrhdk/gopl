package ex01

import "strings"

type WordLineCounter struct {
	Words, Lines int
}

func (c *WordLineCounter) Write(p []byte) (int, error) {
	s := string(p)
	c.Words += len(strings.Fields(s))
	c.Lines += strings.Count(s, "\n") + 1
	return len(p), nil
}
