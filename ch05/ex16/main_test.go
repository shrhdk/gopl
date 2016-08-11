package ex16

import "testing"

func TestJoin(t *testing.T) {
	test(t, "hello world", join(" ", "hello", "world"))
	test(t, "hello!world", join("!", "hello", "world"))
	test(t, "helloworld", join("", "hello", "world"))
	test(t, "hello", join(" ", "hello"))
	test(t, "", join(" "))
}

func test(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
