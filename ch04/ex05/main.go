package ex05

func RemoveDuplication(s []string) []string {
	if len(s) == 0 {
		return s
	}

	out := s[:1]
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			out = append(out, s[i])
		}
	}
	return out
}
