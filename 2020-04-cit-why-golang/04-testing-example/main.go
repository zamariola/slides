package main

import "fmt"

func main() {

	sumResult := sumPositiveValues(2, 3)
	fmt.Println(sumResult)

	sumResult = sumPositiveValues(-2, 3)
	fmt.Println(sumResult)
}

func sumPositiveValues(a, b int) int {

	if a > 0 && b > 0 {
		return a + b
	}
	return 0
}
