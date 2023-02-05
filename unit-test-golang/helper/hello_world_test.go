package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("World")

	// assert.Equal(t, expected, testing function, message error)
	assert.Equal(t, "Hello World", result, "Harusnya Hello World")

	fmt.Println("Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Worlds")

	// assert.Equal(t, expected, testing function, message error)
	require.Equal(t, "Hello World", result, "Harusnya Hello World")

	fmt.Println("Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")

	if result != "Hello World" {
		// Error
		// t.Fail()
		t.Error("Harusnya Hello Worlds")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldSam(t *testing.T) {
	result := HelloWorld("Sam")

	if result != "Hello Sam" {
		// Error
		// t.FailNow()
		t.Fatal("Harusnya Hello Samd")
	}

	fmt.Println("TestHelloWorldSam done")
}
