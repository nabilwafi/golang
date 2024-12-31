package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	}

	return false
}

func TestComparable(t *testing.T) {
	assert.True(t, IsSame[int](1, 1))
	assert.True(t, IsSame[string]("nabil", "nabil"))
	assert.True(t, IsSame[bool](true, true))
}
