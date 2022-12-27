package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	if message != nil {
		fmt.Println("Message Panic:", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERRORRR")
	}
	fmt.Println("Worked")

}

func main() {
	runApp(true)
	fmt.Println("Sam")

}
