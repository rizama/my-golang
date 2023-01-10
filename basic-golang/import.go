package main

import (
	"fmt"
	"my-golang/basic-golang/helper"
)

func main() {
	name := "sam"

	result := helper.SayHello(name)

	fmt.Println(result)
}
