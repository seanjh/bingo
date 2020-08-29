package game

import (
	"log"
)

// Round represents a round of BINGO, including the cage of numbers
// and any cards in play.
type Round struct {
	// ball cage
	cage *Cage
	// cards playing
	cards []*Card
	// winning pattern
	pattern Pattern
	// reference card
	refCard *Card
}

// NewStandardRound returns a standard cage and cards for one round of BINGO.
func NewStandardRound(numCards int) *Round {
	cage := NewStandardCage()
	if numCards < 0 {
		log.Printf("Invalid numCards for Round: %d", numCards)
		numCards = 0
	}
	cards := make([]*Card, 0, numCards)
	for i := 0; i < numCards; i++ {
		cards = append(cards, NewStandardCard())
	}
	refCard := NewStandardCard()
	// TODO: support patterns other than across
	return &Round{cage: cage, cards: cards, pattern: &any{}, refCard: refCard}
}

func (r *Round) hasWinner() bool {
	return true
}

func (r *Round) annouce(pull int) {
	for _, card := range r.cards {
		card.cover(pull)
	}
}

func (r *Round) winners() []*Card {
	winners := make([]*Card, 0, len(r.cards))
	for _, card := range r.cards {
		if card.isWinner() {
			winners = append(winners, card)
		}
	}
	return winners
}

// Simulate returns the winners of the round, immediately running the round to completion.
func (r *Round) Simulate() []*Card {
	winners := r.winners()
	for len(winners) == 0 {
		pull, err := r.cage.take()
		if err != nil {
			log.Panicf("Round cannot proceed: %v", err)
		}
		r.annouce(pull)
		winners = r.winners()
	}
	return winners
}
