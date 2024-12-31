package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		Firstname: "Muhammad",
		Middlename: "Nabil",
		Lastname: "Wafi",
		Hobbies: []string{"Fishing", "Reading", "Gaming"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"Firstname":"Muhammad","Middlename":"Nabil","Lastname":"Wafi","Age":0,"Married":false,"Hobbies":["Fishing","Reading","Gaming"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayEncodeComplex(t *testing.T) {
	cust := Customer{
		Firstname: "Muhammad",
		Middlename: "Nabil",
		Lastname: "Wafi",
		Address: []Address{
			{
				Street: "Jakarta",
				Country: "Indonesia",
				PostalCode: "9999",
			},
			{
				Street: "Bogor",
				Country: "Indonesia",
				PostalCode: "9999",
			},
		},
	}

	bytes, _ := json.Marshal(cust)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecodeComplex(t *testing.T) {
	jsonString := `{"Firstname":"Muhammad","Middlename":"Nabil","Lastname":"Wafi","Age":0,"Married":false,"Hobbies":null,"Address":[{"Street":"Jakarta","Country":"Indonesia","PostalCode":"9999"},{"Street":"Bogor","Country":"Indonesia","PostalCode":"9999"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Address)
}

func TestJSONOnlyArrayDecodeComplex(t *testing.T) {
	jsonString := `[{"Street":"Jakarta","Country":"Indonesia","PostalCode":"9999"},{"Street":"Bogor","Country":"Indonesia","PostalCode":"9999"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestJSONOnlyArrayEncodeComplext(t *testing.T) {
	addresses := []Address{
			{
				Street: "Jakarta",
				Country: "Indonesia",
				PostalCode: "9999",
			},
			{
				Street: "Bogor",
				Country: "Indonesia",
				PostalCode: "9999",
			},
		}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}