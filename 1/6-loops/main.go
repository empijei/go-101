package main

import (
	"fmt"
	"slices"
)

// todo return numbers in order

func main() {
	var data []int
	for {
		var i int
		_, err := fmt.Scanln(&i)
		if err != nil {
			break
		}
		data = append(data, i)
	}
	slices.Sort(data)
	for i, v := range data {
		fmt.Printf("data[%v]: %v\n", i, v)
	}
}
