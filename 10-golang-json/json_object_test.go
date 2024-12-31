package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street string
	Country string
	PostalCode string
}

type Customer struct {
	Firstname string
	Middlename string
	Lastname string
	Age int
	Married bool
	Hobbies []string
	Address []Address
}

func TestJSONObject(t *testing.T) {
	cust := Customer{
		Firstname: "Muhammad",
		Middlename: "Nabil",
		Lastname: "Wafi",
		Age: 22,
		Married: false,
	}

	bytes, _ := json.Marshal(cust)
	
	fmt.Println(string(bytes))
}