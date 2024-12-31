package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello,", firstName, lastName)
}

func main() {
	sayHelloTo("Nabil", "Wafi")
	sayHelloTo("Rizki", "Dani")
	sayHelloTo("Zaki", "Fahmi")
}