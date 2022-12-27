package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 4 {
			continue
		}

		fmt.Println(i)

		if i == 8 {
			break
		}
	}
}
