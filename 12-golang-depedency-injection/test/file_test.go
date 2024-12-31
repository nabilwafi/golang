package test

import (
	"nabilwafi/golang_depedency_injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databases")
	assert.NotNil(t, connection)

	cleanup()
}