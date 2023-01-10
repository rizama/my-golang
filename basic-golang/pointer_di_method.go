package main

import "fmt"

type Man struct {
	Name string
}

// jika ingin menggunakan pointer di struct function, tipe data harus menggunakan "*"
// Jika membuat struct function (struct yang memiliki method) disarankan memakai pointer agar lebih hemat memory
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("*di method", man.Name)
}

func main() {
	sam := Man{"Sem"}
	sam.Married()

	fmt.Println(sam.Name)
}
