package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleWare struct {
	Handler http.Handler
}

func (middleware *LogMiddleWare) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

// Create Error Handler Struct
type ErrorHandler struct {
	Handler http.Handler
}

// Implement Error Handler
func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error Gan")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()

	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handle executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic executed")
		panic("Ups Panic gan")
	})

	logMiddleWare := &LogMiddleWare{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleWare,
	}

	// server -> errorHandler -> logMiddlerware -> mux
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
