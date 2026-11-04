package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		panic(err)
	}
	result := multiply(a, b)
	fmt.Println(result)
}

func multiply(a, b int) int {
	return a * b
}
