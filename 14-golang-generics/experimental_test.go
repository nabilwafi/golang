package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func GetMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	}

	return second
}

func TestGetMin(t *testing.T) {
	assert.Equal(t, 100, GetMin(100, 200))
	assert.Equal(t, 100.0, GetMin(100.0, 200.0))
}

func TestMapGenerics(t *testing.T) {
	data1 := map[string]string{
		"Name": "Nabil",
	}

	data2 := map[string]string{
		"Name": "Nabil",
	}

	assert.True(t, maps.Equal(data1, data2))
}

func TestSliceGenerics(t *testing.T) {
	data1 := []string{
		"nabil",
	}

	data2 := []string{
		"nabil",
	}

	assert.True(t, slices.Equal(data1, data2))
}
