package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func XssExample(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/xss.gohtml"))
	t.ExecuteTemplate(w, "xss.gohtml", map[string]interface{}{
		"Title": "Go-lang Auto Escape",
		"Body": "<p>Hello</p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8132",
		Handler: http.HandlerFunc(XssExample),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func XssDisableEscape(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/xss.gohtml"))
	t.ExecuteTemplate(w, "xss.gohtml", map[string]interface{}{
		"Title": "Go-lang Auto Escape",
		"Body": template.HTML("<p>Hello</p>"),
	})
}

func TestXSSExample(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	XssDisableEscape(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func XssDisableEscapeServe(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/xss.gohtml"))
	t.ExecuteTemplate(w, "xss.gohtml", map[string]interface{}{
		"Title": "Go-lang Auto Escape",
		"Body": template.HTML("<h1>Hello</h1>"),
	})
}

func TestTemplateAutoEscapeServe(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8132",
		Handler: http.HandlerFunc(XssDisableEscapeServe),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}