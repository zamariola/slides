package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := sumPositiveValues(-1, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)
}

func sumPositiveValues(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("Cannot sum nonpositive values")
	}
	return a + b, nil
}
