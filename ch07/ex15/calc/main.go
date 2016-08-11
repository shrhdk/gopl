package main

import (
	"ch07/ex13/eval"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arg := strings.Join(os.Args[1:], "")

	expr, err := eval.Parse(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse Error: %v", err)
		os.Exit(1)
	}

	vars := make(map[eval.Var]bool)
	expr.Check(vars)

	env := make(eval.Env)
	for v, _ := range vars {
		var s string
		fmt.Printf("%s=", v)
		fmt.Scanln(&s)

		f, err := strconv.ParseFloat(s, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid Value: %v", err)
			os.Exit(1)
		}

		env[v] = f
	}

	result := expr.Eval(env)

	fmt.Printf("%s = %f\n", arg, result)
}
