package main

import "fmt"

func sumAll(angka int, numbers ...int) int {
	fmt.Println(numbers)
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)
}
