package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}

	result := nilai / pembagi

	return result, nil
}

func main() {
	result, err := pembagian(100, 0)

	if err != nil {
		fmt.Println("Error Message:", err.Error())
	}else {
		fmt.Println(result)
	}

}