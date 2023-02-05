package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before")

	m.Run() // run all test in hello_world_test.go

	fmt.Println("After")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("World")

	// assert.Equal(t, expected, testing function, message error)
	assert.Equal(t, "Hello World", result, "Harusnya Hello World")

	fmt.Println("Done")
}

func TestSkip(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux")
	}
	result := HelloWorld("World")

	// assert.Equal(t, expected, testing function, message error)
	assert.Equal(t, "Hello World", result, "Harusnya Hello World")

	fmt.Println("Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("World")

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

func TestSubTest(t *testing.T) {
	t.Run("Sam", func(t *testing.T) {
		result := HelloWorld("Sam")
		assert.Equal(t, "Hello Sam", result, "Must Be Hello Sam")
	})

	t.Run("Rizky", func(t *testing.T) {
		result := HelloWorld("Rizky")
		assert.Equal(t, "Hello Rizky", result, "Must Be Hello Rizky")
	})
}

func TestHelloWorldTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld Sam",
			request:  "Sam",
			expected: "Hello Sam",
		},
		{
			name:     "HelloWorld Bro",
			request:  "Bro",
			expected: "Hello Bro",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Harusnya "+test.expected)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Sam")
	}
}

func BenchmarkHelloWorldWaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Waw")
	}
}
