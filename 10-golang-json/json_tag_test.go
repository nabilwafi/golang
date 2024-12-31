package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id: "P001",
		Name: "Apple Mac Book Pro",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	productString := `{"id":"P001","name":"Apple Mac Book Pro","Image_url":"http://example.com/image.png"}`
	productByte := []byte(productString)

	product := &Product{}

	err := json.Unmarshal(productByte, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}