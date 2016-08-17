package palindrome

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func randomNonPalindrome(rng *rand.Rand) string {
	for {
		var runes []rune

		n := rng.Intn(25) // random length up to 24
		if n == 0 {
			continue
		}

		runes = make([]rune, n)

		for i := 0; i < n; i++ {
			runes[i] = rune(rng.Intn(0x100)) // random rune up to \u99
		}

		// simple exam
		i := rng.Intn(n)
		if runes[i] != runes[n-1-i] {
			return string(runes)
		}
	}
}

func TestRandomNoisyPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	n := 0
	for i := 0; i < 100; i++ {
		p := randomNonPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsNoisyPalindrome(%q) = false", p)
			n++
		}
	}
	fmt.Fprintf(os.Stderr, "fail = %d\n", n)
}
