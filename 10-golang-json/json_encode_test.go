package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJSON(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	LogJSON("Nabil")
	LogJSON(1)
	LogJSON(true)
	LogJSON([]string{"Muhammad", "Nabil", "Wafi"})
}