package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONDecode(t *testing.T) {
	jsonString := `{"Firstname":"Muhammad","Middlename":"Nabil","Lastname":"Wafi","Age":22,"Married":false}`
	jsonBytes := []byte(jsonString)

	cust := &Customer{}

	err := json.Unmarshal(jsonBytes, cust)
	if err != nil {
		panic(err)
	}

	fmt.Println(cust)
	fmt.Println(cust.Firstname)
	fmt.Println(cust.Middlename)
	fmt.Println(cust.Lastname)
}