package api

import (
	"encoding/json"
	"log"
	"net/http"
	//
	//"github.com/gorilla/mux"
	//
	//"github.com/seanjh/bingo/backend/game"
)

type createGame struct {
	Id       string `json:"id"`
	ClientId string `json:"clientId"`
}

type example struct{}

func (x *example) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}

func CreateGame(w http.ResponseWriter, req *http.Request) {
	log.Println("creating new game")

	g := createGame{Id: "foo", ClientId: "baz"}
	content, err := json.Marshal(g)
	if err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
		return
	}

	// create new round -- round := game.NewStandardRound(1)
	// create new client/player
	// get client/player GUID
	// get game GUID
	// add game to event bus by game ID
	// return round & client ID

	// r := Round{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(content)
}
