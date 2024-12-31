package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./layouts/header.gohtml", "./layouts/footer.gohtml", "./layouts/index.gohtml"))

	t.ExecuteTemplate(w, "content", map[string]interface{}{
		"Title": "Template Layout",
		"Name": "Nabil",
	})
}

func TestTemplateLayout(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	TemplateLayout(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}