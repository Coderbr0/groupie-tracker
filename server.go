package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
