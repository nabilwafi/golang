package main

import "fmt"

func hello() interface{} {
	return "OK"
}

func main() {
	testing := hello()
	testingString := testing.(string)
	fmt.Println(testingString)
	
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Pesan Error:", err)
		}else {
			fmt.Println("Sistem Berjalan Dengan Baik")
		}
	}()

	// Cara Terbaik
	switch testing := hello(); testing.(type) {
	case string:
		fmt.Println("Ini Adalah String")
	case int:
		fmt.Println("Ini adalah Int")
	default:
		fmt.Println("Tipe data tidak diketahui")
	}

	// ERROR TANPA SELEKSI
	testingInt := testing.(int)
	fmt.Println(testingInt)

}