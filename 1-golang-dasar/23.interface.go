package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}


func main() {

	var nabil Person
	nabil.Name = "Nabil Wafi"

	sayHello(nabil)
}