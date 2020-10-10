package api

import (
	"github.com/gorilla/mux"
)

func NewApiRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	b := NewBroker()
	s.HandleFunc("/subscribe/{clientId}", subscribe(b)).Methods("GET")

	s.HandleFunc("/game", CreateGame).Methods("POST")

	return s
}
