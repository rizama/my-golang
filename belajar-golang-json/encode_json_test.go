package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncodeJson(t *testing.T) {
	LogJson("SAM")
	LogJson(1)
	LogJson(true)
	LogJson([]interface{}{"Rizky", "Sam", "Pratama", 1})
}
