package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./assets/ok.html")
	} else {
		http.ServeFile(w, r, "./assets/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8132",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed assets/ok.html
var okfile string

//go:embed assets/notfound.html
var notfound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, okfile)
	} else {
		fmt.Fprint(w, notfound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8132",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}