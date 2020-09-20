package main

import "fmt"

/*
Struct
- sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
- Struct biasanya representasi data dalam program aplikasi yang kita buat
- data di Struct disimpan didalam field
- sederhananya struct adalah kumpulan dari field

Membuat data Struct
- Struct adalah template data ataupun prototype data
- Struct tidak bisa langsung digunakan
- Kita bisa buat data/object dari struct yang telah kita buat

Struct Literal

*/

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var sam Customer
	sam.Name = "Sam Pratama"
	sam.Address = "Bandung"
	sam.Age = 20
	fmt.Println(sam)

	// Struct Literal
	rizky := Customer{
		Name:    "rizky",
		Address: "Bandung",
		Age:     20,
	}

	fmt.Println(rizky)
}
