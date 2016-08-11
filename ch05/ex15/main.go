package ex15

func min(args ...int) int {
	if len(args) == 0 {
		panic("")
	}

	m := args[0]
	for _, v := range args[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func max(args ...int) int {
	if len(args) == 0 {
		panic("")
	}

	m := args[0]
	for _, v := range args[1:] {
		if m < v {
			m = v
		}
	}
	return m
}

func min2(m int, args ...int) int {
	for _, v := range args {
		if v < m {
			m = v
		}
	}
	return m
}

func max2(m int, args ...int) int {
	for _, v := range args {
		if m < v {
			m = v
		}
	}
	return m
}
