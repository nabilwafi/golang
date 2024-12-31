package main

import "fmt"

func sayHelloTo(name string) string {
	return "Hello, " + name
}

func main() {
	welcome := sayHelloTo("Nabil Wafi")

	fmt.Println(welcome)
}