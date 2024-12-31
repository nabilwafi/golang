package main_test

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("json/customerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		Firstname: "Nabil",
		Lastname: "Wafi",
	}

	encoder.Encode(customer)
}