package main

import "fmt"

func main() {
	var in string
	_, err := fmt.Scanln(&in)
	if err != nil {
		panic(err)
	}
	fmt.Println(in)
}
