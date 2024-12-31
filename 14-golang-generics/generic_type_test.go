package belajar_golang_generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, v := range bag {
		fmt.Println(v)
	}
}

func TestPrintBag(t *testing.T) {
	PrintBag(Bag[int]{1, 2, 3})
	PrintBag(Bag[string]{"Muhammad", "Nabil", "Wafi"})
}
