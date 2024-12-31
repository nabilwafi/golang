package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Tipe Array
	var array[5]int

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	fmt.Println(array)

	var array2 = [5]int{1,2,3,4,5}
	fmt.Println(array2)

	// Tipe Slice
	var fruits = []string{
		"apple",
		"orange",
		"mango",
		"papaya",
		"banana",
	}
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	var newFruits = fruits[3:]

	fmt.Println(newFruits)
	
	newFruits[0] = "Tomato"
	var newFruits2 = append(newFruits, "jaka")

	fmt.Println(newFruits2)
	fmt.Println(newFruits)
	fmt.Println(fruits)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "January" 
	newSlice[1] = "February" 
	fmt.Println("Alamat memori sebelum append:", &newSlice)
	fmt.Println(reflect.ValueOf(newSlice).Pointer())
	
	newSlice = append(newSlice, "March", "April", "May")
	
	fmt.Println("Alamat memori setelah append pertama:", &newSlice)
	fmt.Println(reflect.ValueOf(newSlice).Pointer())
	
	// Menambahkan elemen yang melebihi kapasitas
	newSlice = append(newSlice, "June", "July")
	fmt.Println("Alamat memori setelah append pertama:", &newSlice)
	fmt.Println(reflect.ValueOf(newSlice).Pointer())

	// Tipe Map
	var newMap = make(map[string]string)
	newMap["name"] = "Muhammad Nabil Wafi"
	newMap["umur"] = "12"
	
	fmt.Println(newMap)

	newMap2 := map[string]string{
		"name_car": "Mazda",
		"spec": "Mazda2",
	}

	fmt.Println(newMap2)
}