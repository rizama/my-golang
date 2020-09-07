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

	// months[5] = "Sam"
	// fmt.Println(slice1)

	// slice1[0] = "Pratama"
	// fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Rizky")
	fmt.Println(slice3)
	slice3[1] = "Bukan_desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "MSI"
	newSlice[1] = "ASUS"
	var hehe = append(newSlice, "Sam")

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(hehe)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	var iniArray = [...]int{1, 2, 3, 4, 5}
	var iniSlice = []int{1, 2, 3, 4, 5}

}
