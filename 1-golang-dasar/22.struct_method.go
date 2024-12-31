package main

import "fmt"

type customer struct {
	Name    string
	Age     int
	Address string
}

func (a customer) sayHello() {
	fmt.Println("Hello,", a.Name)
}

func main() {

	customer := customer{
		Name: "Nabil Wafi",
		Age: 15,
		Address: "Jl. FAFFUUUU",
	}

	customer.sayHello()
}