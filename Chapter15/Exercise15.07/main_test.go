package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func Test_posted(t *testing.T) {
	hdl, err := NewHello("./index.html")
	if err != nil {
		t.Error(err)
	}
	srv := httptest.NewServer(hdl)
	form := url.Values{}
	form.Add("name", "john")
	form.Add("surname", "smith")
	form.Add("age", "")
	req, err := http.NewRequest("POST", srv.URL, strings.NewReader(form.Encode()))
	if err != nil {
		t.Error(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := srv.Client().Do(req)
	if err != nil {
		t.Error(err)
	}

	expected, err := os.ReadFile("./teststatics/filled.html")
	if err != nil {
		t.Error(err)
	}

	actual := make([]byte, rsp.ContentLength)
	_, err = rsp.Body.Read(actual)
	if err != io.EOF {
		t.Error(err)
	}

	if string(actual) != string(expected) {
		t.Errorf("\n%s\n%s", string(expected), string(actual))
	}
}
