package server

import (
	// "fmt"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}


/*
1.
package main

import (
        "net/http"
)

func main() {
        http.Handle("/", http.FileServer(http.Dir("./static")))
        http.ListenAndServe(":3000", nil)
}

2.
package main

import (
    "log"
    "net/http"
)

func main() {
    log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

3.
func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

https://blog.logrocket.com/creating-a-web-server-with-golang/
*/