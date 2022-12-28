package main

import "fmt"

// interface yang tidak memiliki method kontrak satupun,
// hal ini membuat secara otomatis tipe data akan menjadi implementasinya
// dengan mnenggunakan interface kosong, kita bisa memasukan tipe data apapun kedalamnya
// bisa dikatakan sebagai tipe 'any' jika dalam typescript

// fungsi ini menerima parameter tipe int, dan 'any'. serta dengan tipe return 'any'
func Freedom(i int, apapun interface{}) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return apapun
	}
}

func main() {
	data := Freedom(5, "sam")

	fmt.Println(data)
}
