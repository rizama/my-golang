package main

import "fmt"

func sayGoodby(name string) string {
	return "GoodBye " + name
}

func main() {
	goodbye := sayGoodby
	fmt.Println(goodbye("Sam"))
}
