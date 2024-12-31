package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func redirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Redirect")
}

func redirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-to", http.StatusPermanentRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", redirectTo)
	mux.HandleFunc("/redirect-from", redirectFrom)

	server := http.Server{
		Addr: "localhost:8132",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
