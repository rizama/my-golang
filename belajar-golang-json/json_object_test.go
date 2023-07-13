package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Costumer struct {
	Firstname  string
	Middlename string
	Lastname   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONObject(t *testing.T) {
	constumer := Costumer{
		Firstname:  "Rizky",
		Middlename: "Sam",
		Lastname:   "Pratama",
		Age:        27,
		Married:    false,
	}

	bytes, err := json.Marshal(constumer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
