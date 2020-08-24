package main

import "fmt"

func main() {
	name := "sams"

	if name == "sam" {
		fmt.Println("Hello", name)
	} else if name == "sams" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Who are you?")
	}

	// shot statement
	if length := len(name); length < 5 {
		fmt.Println("Haha")
	}

}
