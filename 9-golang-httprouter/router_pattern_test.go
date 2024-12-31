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

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Product ", p.ByName("id"), " Item ", p.ByName("itemId"))
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8132/products/1/items/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestRouterCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Image: ", p.ByName("image"))
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8132/images/small/test.png", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	assert.Equal(t, "Image: /small/test.png", string(body))
}