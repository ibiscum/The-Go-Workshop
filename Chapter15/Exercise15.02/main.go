package main

import (
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/chapter1", func(w http.ResponseWriter, r *http.Request) {
		msg := "Chapter 1"
		_, err := w.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}
