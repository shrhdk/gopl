package ex11

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	sign, num := separateSign(s)
	integer, decimal := separateDecimal(num)
	integer = comma(integer)

	if len(decimal) == 0 {
		return sign + integer
	}

	return sign + integer + "." + decimal
}

func separateDecimal(s string) (string, string) {
	if strings.Contains(s, ".") {
		sp := strings.Split(s, ".")
		return sp[0], sp[1]
	}

	return s, ""
}

func separateSign(s string) (string, string) {
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		return s[0:1], s[1:]
	}

	return "", s
}

func comma(s string) string {
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
