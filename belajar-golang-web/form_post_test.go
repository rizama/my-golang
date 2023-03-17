package belajar_golang_web

import (
	f "fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	first := request.PostForm.Get("first") // cara 1, harus diparsing dulu
	last := request.PostFormValue("last")  // cara 2, tanpa parsing

	f.Fprintf(writer, "Hello %s %s", first, last)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first=sam&last=pratama")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081/posts", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	f.Println(bodyString)
}
