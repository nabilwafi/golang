package main

import "fmt"

func testing() interface{} {
	return true
}

func main() {
	var a = testing()

	fmt.Println(a)
}