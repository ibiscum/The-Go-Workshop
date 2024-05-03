package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageWithCounter struct {
	counter int
	heading string
	content string
}

func (h *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.counter++
	msg := fmt.Sprintf("<h1>%s</h1>\n<p>%s<p>\n<p>Views: %d</p>", h.heading, h.content, h.counter)
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	hello := PageWithCounter{heading: "Hello World", content: "This is the main page"}
	cha1 := PageWithCounter{heading: "Chapter 1", content: "This is the first chapter"}
	cha2 := PageWithCounter{heading: "Chapter 2", content: "This is the second chapter"}

	http.Handle("/", &hello)
	http.Handle("/chapter1", &cha1)
	http.Handle("/chapter2", &cha2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
