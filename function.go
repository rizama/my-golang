package main

import "fmt"

func saySomething() {
	fmt.Println("Sams")
}

func sayMyName(name string) {
	fmt.Println(name)
}

func main() {
	saySomething()
	sayMyName("Rizky")
}
