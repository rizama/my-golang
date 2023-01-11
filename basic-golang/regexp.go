package main

import (
	"fmt"
	"regexp"
)

/*
Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
https://github.com/google/re2/wiki/Syntax
https://golang.org/pkg/regexp/
*/

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eko eka edo eto eyo eki", -1)
	fmt.Println(search)
}
