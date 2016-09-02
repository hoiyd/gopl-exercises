package main

import (
	"./eval"
	"fmt"
	// "math"
	"os"
)

func main() {
	expr, err := parseAndCheck(os.Args[1])
	if err != nil {
		fmt.Printf("bad expr: %s\n" + err.Error())
		return
	}
	// result := expr.Eval(eval.Env{"x": 10, "y": 5, "r": 20, "pi": math.Pi})
	fmt.Println(expr)
}

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" && v != "pi" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}
