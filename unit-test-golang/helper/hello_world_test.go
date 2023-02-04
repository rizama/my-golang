package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")

	if result != "Hello Worlds" {
		// Error
		// t.Fail()
		t.Error("Harusnya Hello Worlds")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldSam(t *testing.T) {
	result := HelloWorld("Sam")

	if result != "Hello Samd" {
		// Error
		// t.FailNow()
		t.Fatal("Harusnya Hello Samd")
	}

	fmt.Println("TestHelloWorldSam done")
}
