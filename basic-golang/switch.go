package main

import "fmt"

func main() {
	switch name := "jos"; name {
	case "eko":
		fmt.Println(name)
	case "sam":
		fmt.Println(name)
	case "sams":
		fmt.Println(name)
	default:
		fmt.Println("no result")
	}
}
