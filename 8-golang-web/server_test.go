package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8800",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr: "localhost:8800",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World Images")
	})

	mux.HandleFunc("/images/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumbnail")
	})

	server := http.Server{
		Addr: "localhost:8800",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:8800",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func TestHttp(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8800/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func QueryParam(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Printf("no name")
		}else {
		fmt.Printf("hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8800/hello?name=Nabil", nil)
	recorder := httptest.NewRecorder()

	QueryParam(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func MultiQueryParam(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	names := q["name"]

	fmt.Fprint(w, strings.Join(names, ""))
}

func TestMultiQueryParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8800/hello?name=Nabil&name=Eko", nil)
	recorder := httptest.NewRecorder()

	MultiQueryParam(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-powered-by", "Nabil Wafi")
	fmt.Println(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, req)

	res := recorder.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println(res.Header.Get("x-powered-by"))
}



