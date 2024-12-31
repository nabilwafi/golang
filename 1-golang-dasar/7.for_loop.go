package main

import "fmt"

func main() {
	
	// FOR Pertama
	counter := 0;
	for counter < 10 {
		fmt.Println("Counter ke-",counter)
		counter++
	}

	fmt.Println()
	// FOR KEDUA
	for counter := 0; counter < 10; counter++ {
		fmt.Println("Counter ke-",counter)
	}
	
	fmt.Println()
	// FOR KETIGA
	name := []string{"Nabil", "Faris", "Restu", "Rifyal"}
	for index, value := range name {
		fmt.Println("index ke", index, "dengan isi", value)
	}

	fmt.Println()
	for counter := 0; counter < len(name); counter++ {
		fmt.Println(name[counter])
	}

	fmt.Println()
	mapping := map[string]string{
		"name": "Muhammad Nabil Wafi",
		"age": "22",
		"hobby": "Fishing",
	}
	for index, value := range mapping {
		fmt.Println("keynya adalah", index, "dengan isi", value)
	}

}