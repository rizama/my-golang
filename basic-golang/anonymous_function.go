package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("You Are Not Blocked", name)
	}
}

func main() {
	// Anonymous function di simpan ke dalam variable
	blacklist := func(name string) bool {
		return name == "anjay"
	}
	registerUser("anjay", blacklist)

	// Anonymous function di simpan langsung ke parameter
	registerUser("sem", func(name string) bool {
		return name == "anjay"
	})
}
