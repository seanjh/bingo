package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", base)

	log.Println("starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
