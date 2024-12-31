package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	handler http.Handler
}

func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	m.handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	handler http.Handler
}

func (h *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("RECOVER:", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
		}
	}()
	h.handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Execute")
		fmt.Fprint(w, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter,r *http.Request) {
		fmt.Println("Foo Execture")
		fmt.Fprint(w, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter,r *http.Request) {
		fmt.Println("Panic Execute")
		panic("ups")
	})

	logMiddleware := &LogMiddleware{
		handler: mux,
	}

	errorHandler := &ErrorHandler{
		handler: logMiddleware,
	}

	server := http.Server{
		Addr: "localhost:8132",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}