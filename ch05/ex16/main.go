package ex16

import "bytes"

func join(sep string, a ...string) string {
	var b bytes.Buffer
	for i, s := range a {
		if i != 0 {
			b.WriteString(sep)
		}
		b.WriteString(s)
	}
	return b.String()
}
