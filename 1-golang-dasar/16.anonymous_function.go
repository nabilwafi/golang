package main

import "fmt"

func registerName (name string, blockedName func(string) bool) string {

	isBlocked := blockedName(name)

	if isBlocked {
		return "You are blocked " + name
	}

	return "Welcome, " + name

}

func main() {

	result := registerName("Wafi", func(name string) bool {
		return name == "Nabil"
	})

	blockedName := func (name string) bool  {
		return name == "Wafi"
	}

	result2 := registerName("Wafi", blockedName)

	fmt.Println(result)
	fmt.Println(result2)
}