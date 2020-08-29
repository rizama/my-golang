package main

import "fmt"

func saySomething() {
	fmt.Println("Sams")
}

func sayMyName(name string) {
	fmt.Println(name)
}

func sayMyAge(age int) int {
	return age
}

func main() {
	saySomething()
	sayMyName("Rizky")
	fmt.Println(sayMyAge(24))
}
