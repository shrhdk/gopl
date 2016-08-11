package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestWordFreq(t *testing.T) {
	test(t, "abc abc efg", map[string]int{
		"efg": 1,
		"abc": 2,
	})
}

func test(t *testing.T, given string, expected map[string]int) {
	actual := wordfreq(strings.NewReader(given))
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\ngot  %c\nwant %c", actual, expected)
	}
}
