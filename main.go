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

	fmt.Fprintf(w, "I am an Infrastructure/SRE/DevOps professional\n")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about-me", aboutMeHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
