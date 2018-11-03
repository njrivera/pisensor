package main

import (
	"log"
	"net/http"
)

const (
	port = "5557"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Printf("Running REST API on port %s...", port)
	http.ListenAndServe(":"+port, nil)
}
