package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	// FOR statement
	// Init -> dieksekusi sebelum for berjalan (1kali jalan)
	// Post -> dieksekusi diakhir tiap perulangan (tiap perulangan)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// FOR range
	names := [...]string{"Rizky", "Sam", "Pratama"}
	for _, name := range names {
		fmt.Println(name)
	}
}
