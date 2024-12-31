package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](value T) T {
	fmt.Println(value)
	return value
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Nabil")
	assert.Equal(t, "Nabil", result)

	var resultNumber int = Length[int](80)
	assert.Equal(t, 80, resultNumber)
}
