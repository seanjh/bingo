package main

import (
	"log"

	"github.com/seanjh/bingo/game"
)

func main() {
	log.Print("Starting a new BINGO round")
	round := game.NewStandardRound(1)
	log.Print("Running simulation...")
	winners := round.Simulate()
	log.Printf("Winners: %v\n", winners)
}
