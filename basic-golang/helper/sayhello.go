// dalam 1 package tidak boleh ada fungsi, struct, variable yang sama

package helper

var version = "1.0.0"              // => tidak bisa diakses dari luar Package
var Application = "Belajar Golang" // => bisa diakses dari luar Package

// nama Function Package harus Huruf besar awalnya (Camel Case) agar dapat di akses oleh Package lain.
func SayHello(name string) string {
	return "Hello " + name
}

// tidak bisa diakses dari luar Package, karena huruf awalnya kecil.
func sayGoodBye(name string) string {
	return "Bye " + name
}
