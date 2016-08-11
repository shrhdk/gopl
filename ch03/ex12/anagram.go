package ex12

import (
	"bytes"
	"reflect"
	"unicode"
)

func isAnagram(s1, s2 string) bool {
	s1 = normalize(s1)
	s2 = normalize(s2)

	if s1 == s2 {
		return false
	}

	d1 := analyze(s1)
	d2 := analyze(s2)

	return reflect.DeepEqual(d1, d2)
}

func analyze(s string) map[rune]int {
	d := make(map[rune]int)

	for _, r := range []rune(s) {
		d[r]++
	}

	return d
}

func normalize(s string) string {
	var buf bytes.Buffer

	for _, r := range []rune(s) {
		if unicode.IsSpace(r) {
			continue
		}

		buf.WriteRune(normalizeRune(r))
	}

	return buf.String()
}

func normalizeRune(r rune) rune {
	if unicode.Is(unicode.Katakana, r) {
		return r - 'ア' + 'あ'
	}

	if unicode.IsLetter(r) {
		return unicode.ToUpper(r)
	}

	return r
}
