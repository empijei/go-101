package main

import (
	"errors"
	"fmt"
)

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

	m, err := ComputeMin(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The min value is %v", m)
}

// ComputeMin returns the minimum value of the slice, or an error if the slice is empty.
func ComputeMin(data []int) (int, error) {
	// TODO implement
	return 0, errors.New("slice is empty")
}
