package main

import (
	"fmt"
	"os"
)

/**
Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan  disemua sistem operasi)
https://golang.org/pkg/os/
*/

func main() {

	// argumen yagn di passing dari code run
	args := os.Args
	fmt.Println(args)

	arg1 := args[1]
	fmt.Println(arg1)
	// run: go run os.go arg1 arg2 arg3
	// print: [C:\Users\Sicepat\AppData\Local\Temp\go-build4096914174\b001\exe\os.exe arg1 arg2 arg3]

	// get hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(hostname)

	// set env
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

	// run commnad:
	// 	- export APP_USERNAME=root
	// 	- export APP_PASSWORD=root
	// 	- go run os.go arg1 arg2 arg3
}
