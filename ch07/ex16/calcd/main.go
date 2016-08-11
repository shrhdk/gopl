package main

import (
	"ch07/ex13/eval"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var calc = template.Must(template.New("calc").Parse(`<!DOCTYPE html>
<html>
<head>
	<title>Calc</title>
</head>
<body>
	<h1>Calc</h1>
	<form method="POST">
		<input type="text" name="expr" value="{{.Expr}}" />
		<input type="submit" value="calc" />
	</form>
	<div>{{.Result}}</div>
</body>
</html>`))

type exprAndResult struct {
	Expr   string
	Result string
}

func main() {
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exprs := r.FormValue("expr")

		if exprs == "" {
			calc.Execute(w, exprAndResult{"", ""})
			return
		}

		expr, err := eval.Parse(exprs)

		result := ""
		if err != nil {
			result = fmt.Sprintf("%v", err)
		} else {
			env := make(eval.Env)
			result = fmt.Sprint(expr.Eval(env))
		}

		calc.Execute(w, exprAndResult{exprs, result})
	}))

	sockAddress := "localhost:8000"
	fmt.Println("Serving on " + sockAddress)
	log.Fatal(http.ListenAndServe(sockAddress, nil))
}
