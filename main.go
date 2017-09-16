package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	log.Println("call me on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	log.Printf("%s request from %s", r.Method, r.RemoteAddr)
	io.WriteString(w, "Hello Pimeo!\n")
}
