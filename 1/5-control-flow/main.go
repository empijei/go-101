package main

import "fmt"

func main() {
	var a, b int
	var op string
	_, err := fmt.Scanln(&a, &op, &b)
	if err != nil {
		panic(err)
	}
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
		panic("bad operator: " + op)
	}
	fmt.Println(res)
}

// TODO for and if-else, scoping
