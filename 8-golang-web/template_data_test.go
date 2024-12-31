package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/index.gohtml"))
	t.ExecuteTemplate(w, "index.gohtml", map[string]interface{}{
		"Title": "Template With Map Data",
		"Name": "Nabil",
	})
}

func TestTemplateDataMap(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	TemplateDataMap(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Addr string
} 

type Page struct {
	Title string
	Name string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/index2.gohtml"))
	t.ExecuteTemplate(w, "index2.gohtml", Page{
		Title: "Template With Struct",
		Name: "Nabil",
		Address: Address{
			Addr: "Jl Duar",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	TemplateDataStruct(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}