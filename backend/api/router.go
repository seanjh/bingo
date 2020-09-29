package api

import (
	"github.com/gorilla/mux"
)

func NewApiRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/subscribe", Subscribe)

	s.HandleFunc("/game", CreateGame).Methods("POST")

	return s
}
