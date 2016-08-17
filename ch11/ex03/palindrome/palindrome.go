package palindrome

import "unicode"

func normalize(r rune) rune {
	if unicode.IsLetter(r) && unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}
	return r
}

// IsPalindrome returns true if given string is palindrome.
func IsPalindrome(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		j := len(r) - i - 1

		r1 := normalize(r[i])
		r2 := normalize(r[j])

		if r1 != r2 {
			return false
		}
	}
	return true
}
