package ex07

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	test(t, nil, nil)

	test(t, []byte(""), []byte(""))

	test(t, []byte("あいうえお"), []byte("おえういあ"))
}

func test(t *testing.T, given, expected []byte) {
	Reverse(given)
	if !reflect.DeepEqual(given, expected) {
		t.Errorf("\ngot  %v\nwant %v", string(given), string(expected))
	}
}
