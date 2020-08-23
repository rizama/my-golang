package main

import "fmt"

func main() {
	// cara 1
	var names [3]string
	names[0] = "Rizky"
	names[1] = "Sam"
	names[2] = "Pratama"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// cara 2
	var values = [4]int{
		10, 20, 30, 40,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}
