package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type NumberApp interface {
	~int | int8 | int32 | int64 | float32 | float64
}

func MinApp[T NumberApp](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestApproximation(t *testing.T) {
	assert.Equal(t, Age(100), MinApp[Age](Age(100), Age(200)))
}
