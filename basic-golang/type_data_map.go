package main

import "fmt"

func main() {
	// map[typeKey]typeValue
	// make(map[typeKey]typeValue)
	person := map[string]string{
		"name":    "Sam",
		"address": "Bandung",
	}

	person["gender"] = "Lakilaki"

	fmt.Println(person)
	fmt.Println(len(person))

	// make new map
	books := make(map[string]string)
	books["harry"] = "potter"
	books["one_piece"] = "luffy"

	fmt.Println(books)

	delete(books, "one_piece")

	fmt.Println(books)

}
