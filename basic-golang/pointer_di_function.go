package main

import "fmt"

type Address struct {
	Kota     string
	Provinsi string
	Negara   string
}

// Parameter function tanpa Pointer, sehingga perubahan didalam fungsi ini tidak mempengaruhi original data yang dibawa
func ChangeAddressToIndonesia(address Address) {
	address.Negara = "Indonesia"
}

// Parameter function dengan Pointer (*), sehingga perubahan didalam fungsi ini akan mempengaruhi original data (merubah data sumber yang dibawa)
func ChangeAddressToIndonesiaPointer(address *Address) {
	address.Negara = "Indonesiahh"
}

func main() {
	address := Address{
		Kota:     "Bandung",
		Provinsi: "Jabar",
		Negara:   "Indo",
	}

	fmt.Println("ORIGINAL DATA", address)

	ChangeAddressToIndonesia((address))
	fmt.Println("FUNGSI TANPA POINTER", address) // address ini tidak akan berubah

	// addressPointer := &address
	var addressPointer *Address = &address
	ChangeAddressToIndonesiaPointer((addressPointer))
	fmt.Println("FUNGSI DENGAN POINTER", address) // address akan berubah
}
