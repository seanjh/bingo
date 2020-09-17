package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

type handler = func(http.ResponseWriter, *http.Request)

func logRequestMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", dump)
		handler(w, req)
	}
}

func main() {
	http.HandleFunc("/", logRequestMiddleware(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, foo!\n")
	}))
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// UI
// /match/:id - join confirmation OR game

// API
// /api/v1/match - POST create
// /api/v1/match/:id - GET join confirmation or game
// /api/v1/match/:id/join - POST

// MATCHMAKING - REST
// create new game
// join existing game

// ROUND
// subscribe to game events (calls)
// ~~~ event-driven ~~~
// cover card
// call bingo

// ROUND EVENTS
// call ball
// bingo call
// bingo confirmed
// bingo denied
