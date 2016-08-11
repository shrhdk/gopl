package ex06

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	test(t, nil, nil)

	test(t, []byte(""), []byte(""))

	test(t, []byte("あ"), []byte("あ"))

	test(t, []byte("\t\t\t"), []byte(" "))
	test(t, []byte("\t\t\tあ"), []byte(" あ"))
	test(t, []byte("あ\t\t\t"), []byte("あ "))
	test(t, []byte("あ\t\t\tい"), []byte("あ い"))

	test(t, []byte("   "), []byte(" "))
	test(t, []byte("   あ"), []byte(" あ"))
	test(t, []byte("あ   "), []byte("あ "))
	test(t, []byte("あ   い"), []byte("あ い"))
}

func test(t *testing.T, given, expected []byte) {
	given = CompressSpaces(given)
	if !reflect.DeepEqual(given, expected) {
		t.Errorf("\ngot  %v\nwant %v", string(given), string(expected))
	}
}
