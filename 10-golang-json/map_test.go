package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id": "p001", "name": "Apple Mac Book Pro", "price": 2000}`
	jsonBytes := []byte(jsonString)

	var res map[string]interface{}
	json.Unmarshal(jsonBytes, &res)

	fmt.Println(res)
	fmt.Println(res["id"])
	fmt.Println(res["name"])
	fmt.Println(res["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id": "P0001",
		"name": "Apple Mac Book Pro",
		"price": 2000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}