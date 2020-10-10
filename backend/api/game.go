package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	//
	//"github.com/gorilla/mux"
	//
	//"github.com/seanjh/bingo/backend/game"
)

type Persisted interface {
	Create() error
}

type game struct {
	Id uuid.UUID `json:"id"`
}

func newGame() (*game, error) {
	g := &game{}

	id, err := uuid.NewRandom()
	if err != nil {
		return g, err
	}
	g.Id = id
	return g, nil
}

func CreateGame(w http.ResponseWriter, req *http.Request) {
	log.Println("creating new game")

	g, err := newGame()
	if err != nil {
		http.Error(w, "failed to create game", http.StatusInternalServerError)
		return
	}

	content, err := json.Marshal(g)
	if err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(content)
}
