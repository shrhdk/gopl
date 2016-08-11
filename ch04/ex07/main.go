package ex07

func Reverse(s []byte) {
	rs := []rune(string(s))

	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	copy(s, []byte(string(rs)))
}
