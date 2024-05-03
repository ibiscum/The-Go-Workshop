package main

import (
	"html/template"
	"log"
	"net/http"
)

type Visitor struct {
	Name    string
	Surname string
	Age     string
}

type Hello struct {
	tpl *template.Template
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	vst := Visitor{}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(400)
			return
		}

		vst.Name = r.Form.Get("name")

		vst.Surname = r.Form.Get("surname")
		vst.Age = r.Form.Get("age")
	}

	err := h.tpl.Execute(w, vst)
	if err != nil {
		log.Fatal(err)
	}
}

// NewHello returns a new Hello handler
func NewHello(tplPath string) (*Hello, error) {
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return nil, err
	}
	return &Hello{tmpl}, nil
}

func main() {
	hello, err := NewHello("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", hello)

	http.HandleFunc("/form", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./form.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
