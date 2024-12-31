package main_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware LogMiddleware) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request")
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "GET")
	})

	middleware := LogMiddleware{
		router,
	}

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8132/", nil)
	rec := httptest.NewRecorder()

	middleware.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	
	assert.Equal(t, "GET", string(body))
}