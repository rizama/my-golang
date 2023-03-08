package main_test

import (
	"embed"
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

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var pathFiles embed.FS

func TestMultipleFilesWithPath(t *testing.T) {
	dir, _ := pathFiles.ReadDir("files")
	for _, v := range dir {
		if !v.IsDir() {
			fmt.Println(v.Name())
			content, _ := pathFiles.ReadFile("files/" + v.Name())
			fmt.Println("content:", string(content))
		}
	}
}
