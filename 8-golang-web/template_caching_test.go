package golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var templatess embed.FS

var myTemplates = template.Must(template.ParseFS(templatess, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "index.gohtml", map[string]interface{}{
		"Title": "Template Caching",
		"Name": "Nabil",
	})
}

func TestTemplateCaching(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	TemplateCaching(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}