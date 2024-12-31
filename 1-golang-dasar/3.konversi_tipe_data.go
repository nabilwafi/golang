package main

import (
	"fmt"
)

func main() {
	var num1 int32 = 128
	var conNum1 int64 = int64(num1)
	var conNum2 int8 = int8(num1)

	fmt.Println(num1)
	fmt.Println(conNum1)
	fmt.Println(conNum2)

	var name = "nabil"
	var num2 = name[0]
	var conString = string(num2)

	fmt.Println(name)
	fmt.Println(num2)
	fmt.Println(conString)

	type noKTP string
	var ktpNabil noKTP = "1231231242131231"

	fmt.Println(ktpNabil)
}