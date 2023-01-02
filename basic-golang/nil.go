package main

import "fmt"

// Di Golang, setiap membuat variable dengan tipe data tertentu secara otomatis akan dibuatkan default value
// 		Contoh: int = 0, string = "", boolean = false
// 		Sehingga bukan menjadi null atau undefined
// Namun di Golang terdapat data nil, yaitu data kosong
// nil hanya bisa dipakai dibeberapa tipe data saja
// - interface
// - function
// - map
// - slice
// - pointer
// - channel
// nil biasanya digunakan untuk pengecekan jika data yang kita panggil kosong atau tidak.

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("")

	// Pengecekan menggunakan nil
	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	// Pengecekan tanpa menggunakan nil
	if person["name"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}
