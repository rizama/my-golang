package main

import "fmt"

// Pass	by Value
// di Golang semua variable itu di passing by value, bukan by reference
// artinya, jika kita mengirimkan sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya

// POINTER
// pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
// sederhananya dengan kemampuan Pointer , kita bisa melakukan Pass By Reference
// Operator '&' : untuk membuat sebuah variable dengan nilai pointer ke var lain, kita bisa menggunakan operator '&' diikuti nama variable nya
// Operator '*' : penanda variable tersebut pointer ke variable lain. Selain sebagai penanda, operator '*' juga punya kegunaan lain, yaitu:
// 	- saat kita merubah value di variable pointer, maka yang berubah hanya pada variable tersebut saja
// 	- semua variable yang mengacu ke data yang sama tidak akan berubah
//  - Jika kita ingin mengubah seluruh variable yang mengacu pada data tersebut, kita bisa menggunakan operator '*'

type Address struct {
	Kota     string
	Provinsi string
	Negara   string
}

func main() {
	fmt.Println("OPERATOR &")
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1 // copying value

	address2.Kota = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)

	// menggunakan Pointer, address3 pointer ke address1
	address3 := &address1 // <== menggunakan pointer #1
	address3.Negara = "Konoha"

	fmt.Println(address1) // Negara di address1 akan sama dengan address3
	fmt.Println(address3)

	fmt.Println("----------------------------------------")
	fmt.Println("OPERATOR *")

	// Operator '*'
	address4 := Address{"Garut", "Jawa Barat", "Indonesia"}
	// Awal :
	// 	data1 = address4
	fmt.Println(address4, "address4 sebelum diubah oleh variable lain")

	address5 := &address4 // pointer ke address4
	// Menjadi :
	// 	data1 = address4, address5
	address5.Kota = "Tasik" // ubah field Kota di address5, maka Kota di address4 akan ikut berubah

	fmt.Println(address4, "address4 sesudah diubah oleh variable lain lewat pointer")
	fmt.Println(address5, "update di address5")

	// Assign ulang semua field di variable pointer ke value baru
	// tidak bisa langsung dari struct, tetapi harus berupa pointer juga : cannot use (Address literal) (value of type Address) as *Address value in assignment
	// maka harus diikuti oleh operator '&' pada Address, karena address5 merupakan sebuah variable pointer
	address5 = &Address{"Cianjur", "Jabar", "Indonesia"}
	// Awal :
	// 	data1 = address4, address5 (karena address5 ini pointer)
	// Menjadi:
	// 	data1 = address4
	// 	data2 = address5 (membuat data baru karena assign ulang semua field pada variable)

	fmt.Println(address4, "sekarang address4 tidak akan berubah")
	fmt.Println(address5, "assign ulang semua field di address5")
	fmt.Println("----")

	// disini kita akan membuat data baru (address7) sekaligus membuat siapapun yang mengacu pada address4 akan ikut berpindah ke address5
	address6 := Address{"Yogyakarta", "Jawa Tengah", "Indonesia"}
	// Awal :
	// 	data1 = address6
	fmt.Println(address6, "address6 sebelum diubah oleh variable lain")

	address7 := &address6
	// Menjadi :
	// 	data1 = address6, address7 (karena address7 pointer ke address6)
	address7.Kota = "Semarang"

	// Awal :
	// 	data1 = address6, address7 (karena address7 ini pointer ke address6)
	// Menjadi:
	// 	data1 = kosong
	// 	data2 = address7, address6 (karena menggunakan operator '*', address6 ikut berpindah)
	*address7 = Address{"Black", "Panther", "Wakanda"}

	// kita coba buat variable pointer baru dari address6, seharusnya ikut juga menjadi address7
	address8 := &address6

	fmt.Println(address6, "sekarang address6 akan berubah sesuai address7")
	fmt.Println(address7, "address7 menggunakan point *")
	fmt.Println(address8, "address8 = pengikut address6, yang ikut dipaksa berubah ke address7")

	fmt.Println("----")
	// Function New
	// Sebelumnya membuat pointer menggunakan '&' untuk referensi ke variable tertentu
	// untuk membuat pointer kosong bisa menggunakan function new, artinya tidak ada data awal
	// address9 := new(Address) // deklarasi cara 1
	var address9 *Address = new(Address) // deklarasi cara 2
	fmt.Println(address9)
	address9.Kota = "Malang"
	fmt.Println(address9)
}
