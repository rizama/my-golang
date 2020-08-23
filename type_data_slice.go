package main

import "fmt"

func main() {
	// [...] dipakai jika panjang array tidak diketahui
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:8]
	fmt.Println("slice1: ", slice1)
	fmt.Println("panjang slice1: ", len(slice1))
	fmt.Println("kapasitas slice1: ", cap(slice1))

}
