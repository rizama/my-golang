package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// download file with render
func DownloadFileRender(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}

	http.ServeFile(writer, request, "./resources/"+file)
}

func TestDownloadFileRenderServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFileRender),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// download file with force download
func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}

	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(writer, request, "./resources/"+file)
}

func TestDownloadFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
