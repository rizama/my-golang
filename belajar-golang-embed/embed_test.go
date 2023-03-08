package belajar_golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed files/readme.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}
