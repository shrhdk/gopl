package ex10

import "bytes"

func Comma(s string) string {
	var buf bytes.Buffer

	r := len(s) % 3
	buf.WriteString(s[0:r])
	s = s[r:]

	for len(s) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(",")
		}

		buf.WriteString(s[0:3])
		s = s[3:]
	}

	return buf.String()
}
