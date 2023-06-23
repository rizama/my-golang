package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", map[string]interface{}{
		"Title": "Template Action Condition",
		"Name":  "Sam",
		"Address": map[string]interface{}{
			"Street": "Sadang Serang",
		},
	})
}

func TestTemplateAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://lolcahost:8081", nil)
	recorder := httptest.NewRecorder()

	TemplateAction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionCompare(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/compare.gohtml"))
	t.ExecuteTemplate(writer, "compare.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FirstValue": 60,
	})
}

func TestTemplateActionCompare(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://lolcahost:8081", nil)
	recorder := httptest.NewRecorder()

	TemplateActionCompare(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"game", "watching",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://lolcahost:8081", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Address": map[string]interface{}{
			"Street": "BANDUNG",
			"Number": 10,
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://lolcahost:8081", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
