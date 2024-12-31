package main

import "fmt"

func main() {
	// MATEMATIKA
	num1 := 100

	num1++ // ++, --
	fmt.Println(num1)

	num1 += 1 // -, +, *, /, %
	fmt.Println(num1)

	// Perbandingan
	num2 := 100
	num3 := 200

	perbandingan := num2 == num3 // ==, !=, <=, >=, !=
	fmt.Println(perbandingan)

	// Boolean
	var val = true
	var val2 = true

	boolean := val && val2 // &&, ||, !
	fmt.Println(boolean) 
}