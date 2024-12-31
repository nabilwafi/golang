package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{
		"Title": "Template If",
		"Name": "Nabil",
	})
}

func TestTemplateActionIf(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionIf(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/action.gohtml"))

	t.ExecuteTemplate(w, "action.gohtml", map[string]interface{}{
		"Title": "Template Action Operator",
		"FinalValue": 65,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionOperator(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"Fishing", "Reading", "Coding",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionRange(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name": "Nabil",
		"Address": Address{
			Addr: "Jl. Fafifu",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionWith(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}