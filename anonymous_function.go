package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println(name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "wow"
	}

	registerUser("sam", blacklist)
	registerUser("sem", func(name string) bool {
		return name == "sem"
	})
}
