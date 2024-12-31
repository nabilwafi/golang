package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Nabil",
	})
}

func TestTemplateFunction(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	TemplateFunction(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateGlobalFunc(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Nabil",
	})
}

func TestTemplateGlobalFunc(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	TemplateGlobalFunc(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateCreateGlobalFunc(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func (name string)string {
			return strings.ToUpper(name)
		},
	})
	t = template.Must(t.Parse(`{{upper .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Nabil",
	})
}

func TestTemplateCreateGlobalFunc(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	TemplateCreateGlobalFunc(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateCreateGlobalFuncPipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func (value string)string {
			return "Hello " + value
		},
		"upper": func (name string)string {
			return strings.ToUpper(name)
		},
	})
	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Nabil",
	})
}

func TestTemplateCreateGlobalFuncPipelines(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8132", nil)
	rec := httptest.NewRecorder()

	TemplateCreateGlobalFuncPipelines(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}