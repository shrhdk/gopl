package ex19

func foo() (ret int) {
	defer func() { recover() }()
	ret = 1
	panic(nil)
}
