package ex10

import "testing"

type runesSort []rune

func (x runesSort) Len() int           { return len(x) }
func (x runesSort) Less(i, j int) bool { return x[i] < x[j] }
func (x runesSort) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func Test(t *testing.T) {
	tests := []struct {
		given    string
		expected bool
	}{
		{"ab", false},
		{"", true},
		{"a", true},
		{"aa", true},
		{"aaa", true},
	}

	for _, test := range tests {
		actual := IsPalindrome(runesSort(test.given))
		if actual != test.expected {
			t.Errorf("given: %v\nexpected: %v\nactual: %v", test.given, test.expected, actual)
		}
	}
}
