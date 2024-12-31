package main

import "fmt"

type Customer struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// 1
	var customer1 Customer
	customer1.Name = "Vivian"
	customer1.Age = 15
	customer1.Address = "Jl. Fafifu"

	// 2
	customer2 := Customer{
		Name:    "Nabil Wafi",
		Age:     15,
		Address: "Jl. Kinanti Duar",
	}

	fmt.Println(customer2)
}