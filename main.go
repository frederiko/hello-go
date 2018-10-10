package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloHandler)
	log.Println("hit me on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	host, err := os.Hostname()
	if err != nil {
		http.Error(w, "Could not get hostname", 500)
		return
	}

	ip, err := net.LookupIP(host)
	if err != nil {
		http.Error(w, "Could not get ipaddress", 500)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	log.Printf("%s request from %s", r.Method, r.RemoteAddr)
	fmt.Fprintf(w, "You've hit %s (%s) from %s.\n", host, ip[0], r.RemoteAddr)
}
