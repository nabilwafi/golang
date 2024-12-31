package golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	dir := http.Dir("./assets")
	fileserver := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr: "localhost:8132",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed assets
var assets embed.FS

func TestFileServerWebEmbed(t *testing.T) {
	dir, _ := fs.Sub(assets, "assets")
	fileserver := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr: "localhost:8132",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}