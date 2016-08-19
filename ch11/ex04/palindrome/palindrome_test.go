package palindrome

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
	"unicode"
)

func randomSeparator(rng *rand.Rand) rune {
	return separators[rng.Intn(len(separators))]
}

func randomChar(rng *rand.Rand) rune {
	for {
		r := rune(rng.Intn(0x100)) // random rune up to \u99
		if isSeparator(r) {
			continue
		}
		return r
	}
}

func toUpper(r rune) rune {
	if !unicode.IsLetter(r) || !unicode.IsLower(r) {
		return r
	}

	r1 := r
	r = unicode.ToUpper(r)
	if unicode.ToLower(r) != r1 {
		fmt.Printf("cap? %c %c\n", r1, r)
	}

	return r
}

func insertRandomSeparators(rng *rand.Rand, s string) string {
	src := []rune(s)
	var runes []rune
	for i := 0; i < len(src); {
		if rng.Intn(2) == 0 {
			runes = append(runes, randomSeparator(rng))
		} else {
			runes = append(runes, src[i])
			i++
		}
	}
	return string(runes)
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := randomChar(rng)
		runes[i] = r
		runes[n-1-i] = toUpper(r)
	}
	return insertRandomSeparators(rng, string(runes))
	//return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	n := 0
	for i := 0; i < 100; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
			n++
		}
	}
	fmt.Fprintf(os.Stderr, "fail = %d\n", n)
}
