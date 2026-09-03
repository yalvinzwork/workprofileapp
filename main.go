package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Welcome to my work profile!\n")
}

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about-me" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "I am an information technology professional\n")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about-me", aboutMeHandler)

	log.Println("Starting server on :9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
