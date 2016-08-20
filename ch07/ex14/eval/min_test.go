package eval

import "testing"

func TestMinCheck(t *testing.T) {
	tests := []struct {
		s       string
		n       int // num of vars
		success bool
	}{
		{"min()", -1, false},
		{"min(1)", -1, false},
		{"min(1,2)", 0, true},
		{"min(1,2,3)", 0, true},
		{"min(a,b)", 2, true},
		{"min(a,b,c)", 3, true},
		{"min(a,2,3)", 1, true},
		{"min(1,a,3)", 1, true},
		{"min(1,2,a)", 1, true},
	}

	for _, test := range tests {
		expr, err := Parse(test.s)
		if err != nil {
			t.Errorf("Parse(%v) returns %v", expr, err)
		}

		vars := make(map[Var]bool)
		err = expr.Check(vars)
		if !test.success && err == nil {
			t.Errorf("(%v).Check(vars) returns nil, want any error", expr)
		}

		if test.success && err != nil {
			t.Errorf("(%v).Check(vars) returns %v, want nil", expr, err)
		}

		if test.success && test.n != len(vars) {
			t.Errorf("len(vars) is %d, want %d, after (%v).Check(vars)",
				len(vars), test.n, expr)
		}
	}
}

func TestMinEval(t *testing.T) {
	tests := []struct {
		s    string
		want float64
		env  Env
	}{
		{"min(1,2)", 1, nil},
		{"min(1,2,3)", 1, nil},
		{"min(4,3,2)", 2, nil},
		{"min(a,2,3)", 1, map[Var]float64{"a": 1}},
	}

	for _, test := range tests {
		expr, err := Parse(test.s)
		if err != nil {
			t.Errorf("Parse(%v) returns %v", expr, err)
		}
		if got := expr.Eval(test.env); got != test.want {
			t.Errorf("%v got %f, want %f", expr, got, test.want)
		}
	}
}
