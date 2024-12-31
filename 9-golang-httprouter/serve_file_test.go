package main_test

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed assets
var assets embed.FS

func TestServeFileHello(t *testing.T) {
	router := httprouter.New()
	dir, _ := fs.Sub(assets, "assets")
	router.ServeFiles("/files/*filepath", http.FS(dir))

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8132/files/hello.txt", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	
	assert.Equal(t, "Hello", string(body))
}