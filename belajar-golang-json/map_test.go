package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"Id":"1","Name":"SmartPhone","Price":20000000}`
	jsonBytes := []byte(jsonString)

	var product map[string]interface{}
	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product["Id"])
	fmt.Println(product["Name"])
	fmt.Println((product["Price"]))
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "MACBOOK",
		"price": 20000000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
