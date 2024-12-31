package main

import "fmt"

func sum(numbers ...int) (total int) {

	total = 0
	for _, val := range numbers {
		total += val
	}

	return
}

func main() {
	totalNumbers1 := sum(10,10,10,10)

	fmt.Println(totalNumbers1)

	num := []int{10,20,30,40,50}
	total := sum(num...)

	fmt.Println(total)
}