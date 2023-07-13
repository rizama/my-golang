package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	reader, _ := os.Create("customer_output.json")
	decoder := json.NewEncoder(reader)

	customer := Costumer{
		Firstname:  "Sam",
		Middlename: "Pratama",
		Lastname:   "Who",
		Age:        27,
		Married:    false,
		Addresses: []Address{
			{Street: "Coblong", Country: "Indonesia", PostalCode: "40133"},
			{
				Street:     "Sadang Serang",
				Country:    "Indonesia",
				PostalCode: "40133",
			},
		},
	}

	_ = decoder.Encode(customer)

	fmt.Println(customer)
}
