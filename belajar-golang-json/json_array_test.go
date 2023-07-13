package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Costumer{
		Firstname:  "sam",
		Middlename: "sung",
		Lastname:   "terbaik",
		Age:        24,
		Married:    true,
		Hobbies:    []string{"gaming", "computing"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONDecode(t *testing.T) {
	stringJson := `{"Firstname":"sam","Middlename":"sung","Lastname":"terbaik","Age":24,"Married":true,"Hobbies":["gaming","computing"]}`
	jsonBytes := []byte(stringJson)

	customer := &Costumer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Costumer{
		Firstname:  "sam",
		Middlename: "pratamaa",
		Addresses: []Address{
			{
				Street:     "Coblong",
				Country:    "Indonesia",
				PostalCode: "40133",
			},
			{
				Street:     "Sadang Serang",
				Country:    "Indonesia",
				PostalCode: "40133",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONDecodeComplex(t *testing.T) {
	stringJson := `{"Firstname":"sam","Middlename":"","Lastname":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Coblong","Country":"Indonesia","PostalCode":"40133"},{"Street":"Sadang Serang","Country":"Indonesia","PostalCode":"40133"}]}`
	jsonBytes := []byte(stringJson)

	customer := &Costumer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONDecodeComplex(t *testing.T) {
	stringJson := `[{"Street":"Coblong","Country":"Indonesia","PostalCode":"40133"},{"Street":"Sadang Serang","Country":"Indonesia","PostalCode":"40133"}]`
	jsonBytes := []byte(stringJson)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
