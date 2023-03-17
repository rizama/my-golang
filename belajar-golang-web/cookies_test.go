package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookies(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)

	cookie.Name = "X-SAM-NAME"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success create Cokies")
}

func GetCookies(writer http.ResponseWriter, request *http.Request) {
	cookies := request.Cookies()
	fmt.Println(cookies)

	cookie, err := request.Cookie("X-SAM-NAME")
	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookies)
	mux.HandleFunc("/get-cookie", GetCookies)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/?name=sam", nil)
	recorder := httptest.NewRecorder()

	SetCookies(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/", nil)

	cookie := new(http.Cookie)
	cookie.Name = "X-SAM-NAME"
	cookie.Value = "YUHUUU"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()
	GetCookies(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
