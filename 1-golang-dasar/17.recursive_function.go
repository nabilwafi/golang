package main

import "fmt"

func factorial(val int) int {
	if(val == 1) {
		return 1
	}

	return val * factorial(val - 1)
}

func main() {

	result := factorial(5)

	fmt.Println(result)
	fmt.Println(1 * 2 * 3 * 4 * 5)
}