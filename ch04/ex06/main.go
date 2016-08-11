package ex06

import "unicode"

// TODO: use utf8.DecodeRuneInString
func CompressSpaces(s []byte) []byte {
	if len(s) == 0 {
		return s
	}

	rs := []rune(string(s))
	for i, _ := range rs {
		if unicode.IsSpace(rs[i]) {
			rs[i] = ' '
		}
	}

	out := rs[:1]
	for i := 1; i < len(rs); i++ {
		if rs[i-1] != rs[i] {
			out = append(out, rs[i])
		}
	}

	return []byte(string(out))
}
