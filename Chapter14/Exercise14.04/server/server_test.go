package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFileUploadServeHTTP(t *testing.T) {
	server := server{}

	fileDataBuffer := bytes.Buffer{}
	multipartWritter := multipart.NewWriter(&fileDataBuffer)

	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}

	formFile, err := multipartWritter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}

	multipartWritter.Close()

	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(fileDataBuffer.Bytes()))

	r.Header.Set("Content-Type", multipartWritter.FormDataContentType())

	w := httptest.NewRecorder()

	server.ServeHTTP(w, r)
	if w.Body.String() != "test.txt Uploaded!" {
		t.Errorf("Expected 'test.txt Uploaded!' string but received: '%s'", w.Body.String())
	}
}
