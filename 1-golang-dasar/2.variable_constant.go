package main

import "fmt"

func main() {
	// INTEGER AND FLOAT
	var num1 int
	num1 = 1
	fmt.Println(num1)

	var num2 float32
	num2 = 1.5
	fmt.Println(num2)

	// STRING
	name := "Muhamad Nabil Wafi"
	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(name[0])
	fmt.Println(name[1])

	const (
		firstName = "Muhammad"
		middleName = "Nabil"
		lastName = "Wafi"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}