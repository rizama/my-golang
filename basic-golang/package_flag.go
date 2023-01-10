package main

import (
	"flag"
	"fmt"
)

/**
Package flag berisikan fungsionalitas untuk memparsing command line argument
https://golang.org/pkg/flag/
*/

func main() {
	// kita ingin memformat argument seperti ini:
	// go run package_flag.go -host=localhost -user=root -password=root

	// (flag, default_value, description/helper)
	// flag harus sesuai dengan yang di tuliskan pada argument
	// flag ini mengembalikan type Pointer string, maka harus disimpan kedalam variable pointer juga
	host := flag.String("host", "localhost", "Put your Host")
	user := flag.String("user", "sam", "Put your User")
	password := flag.String("password", "sam", "Put your Password")

	// parsing flag (wajib)
	flag.Parse()

	// pakai data arg
	fmt.Println("host", *host)
	fmt.Println("user", *user)
	fmt.Println("password", *password)
}
