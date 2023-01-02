package main

import (
	"errors"
	"fmt"
)

// Golang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya yaitu 'error'
// jadi ketika ingin mengeluarkan/membuat error, tidak perlu membuat error sendiri, sudah ada kontrak interface error ini
// Golang sudah menyediakan library untuk membuat helper error secara mudah, terdapat di package error

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Gagal, Pembagian dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	value, err := Pembagian(10, 0)

	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(value)
	}
}
