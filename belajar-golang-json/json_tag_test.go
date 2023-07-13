package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`        // merubah nama field setelah dikonversi
	Name     string `json:"name"`      // merubah nama field setelah dikonversi
	ImageURL string `json:"image_url"` // merubah nama field setelah dikonversi
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "1",
		Name:     "SmartPhone",
		ImageURL: "qwertyu",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"Id":"1","Name":"SmartPhone","ImageURL":"qwertyu"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
