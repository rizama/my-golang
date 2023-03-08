package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed files/images.png
var logo []byte

//go:embed files/*.txt
var pathFiles embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("logo_new_1.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := pathFiles.ReadDir("files")
	for _, v := range dir {
		if !v.IsDir() {
			fmt.Println(v.Name())
			content, _ := pathFiles.ReadFile("files/" + v.Name())
			fmt.Println("content:", string(content))
		}
	}
}
