package ex12

import "testing"

// TestAnagram

func TestAnagram(t *testing.T) {
	areAnagram(t, "anagrams", "ARS MAGNA")
	areAnagram(t, "あなぐらむ", "なら グアム")
}

func TestNotAnagram(t *testing.T) {
	areNotAnagram(t, "hogehoge", "hogehoge")
	areNotAnagram(t, "hoge hoge", "hogehoge")
	areNotAnagram(t, "fugafuga", "hogehoge")
	areNotAnagram(t, "ふがふが", "ほげほげ")
}

func areAnagram(t *testing.T, s1 string, s2 string) {
	if !isAnagram(s1, s2) {
		t.Errorf("false negative:\ns1: %s\ns2: %s", s1, s2)
	}
}

func areNotAnagram(t *testing.T, s1 string, s2 string) {
	if isAnagram(s1, s2) {
		t.Errorf("false positive:\ns1: %s\ns2: %s", s1, s2)
	}
}

// TestNormalize

func TestNormalize(t *testing.T) {
	testNormalize(t, "", "")

	testNormalize(t, " ", "")
	testNormalize(t, "  ", "")

	testNormalize(t, "A Z", "AZ")
	testNormalize(t, "A  Z", "AZ")

	testNormalize(t, " AZ", "AZ")
	testNormalize(t, "  AZ", "AZ")

	testNormalize(t, "AZ ", "AZ")
	testNormalize(t, "AZ  ", "AZ")

	testNormalize(t, "a", "A")
	testNormalize(t, "ア", "あ")

	testNormalize(t, "aア", "Aあ")
	testNormalize(t, "Aア", "Aあ")
}

func testNormalize(t *testing.T, given, expected string) {
	actual := normalize(given)
	if actual != expected {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}

// TestNormalizeRune

func TestNormalizeRune(t *testing.T) {
	testNormalizeRune(t, 'a', 'A')
	testNormalizeRune(t, 'A', 'A')
	testNormalizeRune(t, 'ア', 'あ')
	testNormalizeRune(t, 'あ', 'あ')
}

func testNormalizeRune(t *testing.T, given, expected rune) {
	actual := normalizeRune(given)
	if actual != expected {
		t.Errorf("\ngot  %c\nwant %c", actual, expected)
	}
}
