package main

import "fmt"

func main() {
	counter := 0
	name := "sam"

	increment := func() {
		fmt.Println("Increment")
		name = "budi"
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
