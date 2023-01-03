package main

import "fmt"

// Type Assertions merupakan kemampuan untuk mengubah tipe data menjadi tipe data yang diinginkan
// fitur ini sering digunakan ketika kita bertemu dengan data interface kosong

func random() interface{} {
	return "OK"
}

func main() {
	defer endApp()
	result := random()

	// karena result diatas hasil dari interface kosong, maka tidak ada kepastian tipe datanya berupa apa
	// maka jika kita yakin result diatas akan selalu string, kita bisa ubah tipe data nya menggunakan Type Assertions
	// Ubah type data menggunakan assertions
	resultString := result.(string)
	fmt.Println(resultString) // "OK"

	// Case jika Salah merubah type assertions, memaksa string di konversi ke int
	resultInt := result.(int) // panic
	fmt.Println(resultInt)

	// Saat salah menggunakan type assertions, maka bisa berakibat panic di applikasi
	// jika panic dan tidak ter-recover , maka aplikasi akan berhenti
	// agar lebih aman, sebaiknya kita menggunakan "switch" expression untuk melakukan type assertions

	// ambil Type data dari hasil random diatas
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("boolean", value)
	default:
		fmt.Println("Unknown type", value)
	}

}

// tambahan code
func endApp() {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	if message != nil {
		fmt.Println("Message Panic:", message)
	}
}
