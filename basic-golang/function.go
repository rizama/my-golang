package main

import "fmt"

// function without parameter
func saySomething() {
	fmt.Println("Sams")
}

// function with parameter
func sayMyName(name string) {
	fmt.Println(name)
}

// function with parameter and return
func sayMyAge(age int) int {
	return age
}

// function with parameter and multiple return
func sayMyAddress(province string, city string) (string, string) {
	return province, city
}

// function with parameter and multiple return
func whatMyAddress() (string, string) {
	return "sadang serang", "cikutra"
}

// function with parameter and named return values
func myFullName() (firstname, middlename, lastname string, age int) {
	firstname = "rizky"
	middlename = "sam"
	lastname = "pratama"
	age = 24

	return
}

func main() {
	saySomething()

	sayMyName("Rizky")

	fmt.Println(sayMyAge(24))

	fmt.Println(sayMyAddress("Jawa Barat", "Bandung"))
	provinsi, _ := whatMyAddress()
	fmt.Println(provinsi)

	_firtsname, _middlename, _lastname, _age := myFullName()
	fmt.Println(_firtsname, _middlename, _lastname, _age)

}
