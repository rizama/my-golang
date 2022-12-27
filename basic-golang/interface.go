package main

import "fmt"

/**
Interface:
- tipe data yang abstract, dia tidak memiliki implementasi langsung
- sebuah interface berisikan definisi-definisi method
- biasanya interface digunakan sebagai kontrak
*/
type HasName interface {
	getName() string
}

// fungsi sayHello memiliki kontrak dengan interface HasName
// setiap ada type data yang mengimplementasi HasName maka dapat mengakses fungsi ini
func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

// struct
type Person struct {
	Name string
	Age  int
}

// struct function, dimana fungsinya sama dengan yang ada pada interface
// secara otomatis fungsi ini sudah bagian dari interface HasName (mengimplementasi HasName), dan
// berhak mengakases fungsi sayHello()
func (person Person) getName() string {
	return person.Name
}

// struct function, bukan bagian dari syarat interface HasName
func (person Person) getAge() int {
	return person.Age
}

func main() {

	// buat data Person
	sam := Person{
		Name: "sam",
		Age:  26,
	}

	// akses fungsi yang memiliki kontrak dengan interface
	sayHello(sam)

	// akses fungsi struct biasa yang tidak berkaitan dengan interface
	fmt.Println(sam.getAge())
}
