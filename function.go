package main

import "fmt"

func saySomething() {
	fmt.Println("Sams")
}

func sayMyName(name string) {
	fmt.Println(name)
}

func sayMyAge(age int) int {
	return age
}

func sayMyAddress(province string, city string) (string, string) {
	return province, city
}

func whatMyAddress() (string, string) {
	return "sadang serang", "cikutra"
}

func main() {
	saySomething()
	sayMyName("Rizky")
	fmt.Println(sayMyAge(24))
	fmt.Println(sayMyAddress("Jawa Barat", "Bandung"))
	provinsi, _ := whatMyAddress()
	fmt.Println(provinsi)
}
