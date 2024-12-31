package main

import "fmt"

func nameFiltered(name string, filter func(string) string) string {
	
	filteredName := filter(name)

	return "Hello, " + filteredName
} 

func spamFiltered(name string) string {
	
	filteredName := name
	if name == "Anjing" {
		filteredName = "..."
	}

	return filteredName
}

func main() {
	name := "Nabil Wafi"
	result := nameFiltered(name, spamFiltered)

	fmt.Println(result)
}