package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{
		"name": name,
	}
}

func main() {
	var person map[string]string = newMap("nabil")

	if person == nil {
		fmt.Println("data kosong")
	}else {
		fmt.Println(person)
	}
}