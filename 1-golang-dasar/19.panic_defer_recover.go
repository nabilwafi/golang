package main

import "fmt"

// DEFER
func testDefer() {
	message := recover()

	if message != nil {
		fmt.Println("Pesan Error:", message)
	}

	fmt.Println("Program Selesai")
}

// PANIC
func testPanic(error bool) {
	if error {
		panic("Sistem Error")
	}
}

// Recover
func testRecover() {
	message := recover()

	if message != nil {
		fmt.Println("Terjadi Error", message)
	}

}

func main() {
	defer testDefer()
	fmt.Println("Program Dimulai")

	testPanic(false)

}