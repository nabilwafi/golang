package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML (w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	SimpleHTML(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./assets/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello How Are You?")
}

func TestSimpleHTMLFile(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	SimpleHTMLFile(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFileGlob(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./assets/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello How Are You?")
}

func TestSimpleHTMLFileGlob(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	SimpleHTMLFileGlob(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

//go:embed assets/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "assets/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello")
}

func TestTemplateEmbed(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	TemplateEmbed(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}