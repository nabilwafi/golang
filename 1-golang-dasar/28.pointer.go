package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var val1 Address = Address{
		City: "Bandung",
		Province: "Jawa Barat",
		Country: "Indonesia",
	}

	var val2 *Address = &val1
	val2.City = "Jakarta"

	*val2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(&val1)
	fmt.Println(val2)

	val3 := new(Address)
	val4 := &val3

	fmt.Println(val3)
	fmt.Println(val4)
	val3.City = "Sukabumi"

	fmt.Println(val3)
	fmt.Println(val4)
}