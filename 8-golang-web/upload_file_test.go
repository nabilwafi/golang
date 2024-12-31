package golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./assets/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./assets"))))
	
	server := http.Server{
		Addr: "localhost:8132",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed assets/GRYe2rSaEAA00jj.jpg
var image []byte

func TestUploadForm2(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Nabil Wafi")

	file, _ := writer.CreateFormFile("file", "NewImage.png")
	file.Write(image)
	writer.Close()

	r := httptest.NewRequest(http.MethodPost, "http://localhost:8132/upload", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()

	Upload(rec, r)

	bodyResponse, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(bodyResponse))
}