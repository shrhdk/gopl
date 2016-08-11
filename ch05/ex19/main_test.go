package ex19

import "testing"

func TestFoo(t *testing.T) {
	if foo() != 1 {
		t.Fail()
	}
}
