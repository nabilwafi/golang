package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	server := http.Server{
		Addr: "localhost:8132",
		Handler: router,
	}
	fmt.Println("Server is running!")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}