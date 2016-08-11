package eval

import "testing"

func Test(t *testing.T) {
	tests := []string{
		"2",
		"0.1",
		"A",
		"A*1*1",
		"1+B+1",
		"sqrt(B)",
		"sqrt(B/1)",
	}

	for _, given := range tests {
		expr1, err := Parse(given)

		if err != nil {
			t.Errorf("%v", err)
		}

		expr2, err := Parse(expr1.String())

		if err != nil {
			t.Errorf("%v", err)
		}

		s1 := expr1.String()
		s2 := expr2.String()

		if s1 != s2 {
			t.Errorf("%s != %s", s1, s2)
		}
	}
}
