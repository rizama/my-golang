package main

import (
	"fmt"
	"my-golang/basic-golang/helper"
)

func main() {
	name := "sam"

	result := helper.SayHello(name)

	fmt.Println(result)

	// helper.sayGoodBye(name) // => error, fungsi  sayGoodBye tidak di export
	helper.SayHello(name) // => berfungsi SayHello sudah di export, karena memakai huruf awal besar

	// fmt.Println(helper.version) //=> error, variable version tidak di export
	fmt.Println(helper.Application) //=> variable Application sudah di export

}
