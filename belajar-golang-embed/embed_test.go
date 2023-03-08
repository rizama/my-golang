package belajar_golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed files/readme.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed files/images.png
var logo []byte

func TestByteSlice(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
