package main

import "fmt"

/**
https://gobyexample.com/interfaces
Interface:
- tipe data yang abstract, dia tidak memiliki implementasi langsung
- sebuah interface berisikan definisi-definisi method
- biasanya interface digunakan sebagai kontrak
*/
type HasName interface {
	getName() string
	isMarried() bool
}

// generic function, fungsi umum yang bisa dipakai oleh banyak type
// fungsi sayHello memiliki kontrak dengan interface HasName
// setiap ada type data yang mengimplementasi HasName maka dapat mengakses fungsi ini
func sayHello(hasName HasName) {
	var isMarried string
	if hasName.isMarried() {
		isMarried = "Yes"
	} else {
		isMarried = "No"
	}

	fmt.Println("Hello", hasName.getName(), ", Are you married?", isMarried)
}

// struct
// we’ll implement the interface on Person types.
type Person struct {
	Name      string
	Age       int
	IsMarried bool
}

// To implement an interface in Go, we just need to implement all the methods in the interface
// struct function, dimana fungsinya sama dengan yang ada pada interface
// secara otomatis fungsi ini sudah bagian dari interface HasName (mengimplementasi HasName), dan
// berhak mengakses fungsi sayHello()
func (person Person) getName() string {
	return person.Name
}
func (person Person) isMarried() bool {
	return person.IsMarried
}

// struct function, bukan bagian dari syarat interface HasName
func (person Person) getAge() int {
	return person.Age
}

func main() {

	// buat data Person
	sam := Person{
		Name:      "sam",
		Age:       26,
		IsMarried: false,
	}

	// akses fungsi yang memiliki kontrak dengan interface
	sayHello(sam)

	// akses fungsi struct biasa yang tidak berkaitan dengan interface
	fmt.Println(sam.getAge())
}
