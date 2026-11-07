package main

import (
	"fmt"
)

func main() {
	var a, b int
	var op string
	_, err := fmt.Scanln(&a, &op, &b)
	if err != nil {
		panic(err)
	}
	res, err := Execute(a, b, op)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func Execute(a, b int, op string) (int, error) {
	var res int
	switch op {
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "^":
		res = a
		for range b - 1 {
			res = res * a
		}
	default:
		return 0, fmt.Errorf("unknown operator: %q", op)
	}
	return res, nil
}
