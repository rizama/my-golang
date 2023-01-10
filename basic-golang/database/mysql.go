package database

import "fmt"

/**
Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database
Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init
*/

var connection string

// Fungsi init ini akan selalu di jalankan (pertama kali) secara otomatis
func init() {
	fmt.Println("INIT DI PANGGIL")
	connection = "MYSQL"
}

// ketika fungsi ini dipanggil, maka fungsi init akan jalan terlebih dahulu
func GetConnection() string {
	return connection
}
