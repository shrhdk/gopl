package ex04

func Rotate(s []int, n int) {
	buf := make([]int, len(s))
	copy(buf, s[len(s)-n:])
	copy(buf[n:], s[:len(s)-n])
	copy(s, buf)
	return
}
