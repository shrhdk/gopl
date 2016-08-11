package ex03

import (
	"bytes"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var b bytes.Buffer
	var visitAll func(t *tree)
	visitAll = func(t *tree) {
		b.WriteString(strconv.Itoa(t.value))
		if t.left != nil {
			visitAll(t.left)
		}
		if t.right != nil {
			visitAll(t.right)
		}
	}
	visitAll(t)
	return b.String()
}
