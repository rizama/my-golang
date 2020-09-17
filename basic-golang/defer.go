// Defer merupakan function yang akan selalu di eksekusi
package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp(value int) {
	defer logging()
	fmt.Println("Run Program")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApp(0)
}
