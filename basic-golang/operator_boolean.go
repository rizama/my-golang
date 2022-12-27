package main

import "fmt"

func main() {
	var nilaiakhir = 98
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiakhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}
