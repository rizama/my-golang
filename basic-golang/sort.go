package main

/**
Package sort adalah package yang berisikan utilitas untuk proses pengurutan
Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
https://golang.org/pkg/sort/
*/

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// slice of User
type UserSlice []User

// implement kontrak Sort Interface (Len)
func (value UserSlice) Len() int {
	return len(value)
}

// implement kontrak Sort Interface (Less)
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

// implement kontrak Sort Interface (Swap)
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Eko", 30},
		{"Jokok", 35},
		{"Budi", 31},
		{"Rudi", 25},
	}

	fmt.Println(users)

	// disini UserSlice yang memiliki kontrak dengan Sort Interface, maka users harus masuk kedalam UserSlice
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
