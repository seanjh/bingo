package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/subscribe", subscribe)

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
