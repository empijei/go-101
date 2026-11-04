package main

import "fmt"

func main() {
	var a, b float64
	var op string
	_, err := fmt.Scanln(&a, &op, &b)
	if err != nil {
		panic(err)
	}
	var res float64
	switch op {
	case "*":
		res = a * b
	case "/":
		res = a * b
	case "+":
		res = a * b
	case "-":
		res = a * b
	default:
		panic("bad operator: " + op)
	}
	fmt.Println(res)
}

// TODO for and if, scoping
