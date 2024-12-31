package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresident() string
}

type MyVicePresident struct {
	Name string
}

func (v *MyVicePresident) GetName() string {
	return v.Name
}

func (v *MyVicePresident) GetVicePresident() string {
	return v.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Nabil", GetName[Manager](&MyManager{Name: "Nabil"}))
	assert.Equal(t, "Nabil", GetName[VicePresident](&MyVicePresident{Name: "Nabil"}))
}
