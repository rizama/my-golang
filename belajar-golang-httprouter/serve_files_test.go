package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed sumber
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	mainDirectory, _ := fs.Sub(resources, "sumber/text")

	router.ServeFiles("/files/*filepath", http.FS(mainDirectory))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Mamen", string(body))
}
