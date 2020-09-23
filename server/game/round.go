package game

import (
	"log"
)

// Round represents a round of BINGO, including the cage of numbers
// and any cards in play.
type Round struct {
	cage    *Cage   // ball cage
	cards   []*Card // cards playing
	pattern Pattern // winning pattern
	refCard *Card   // reference card
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

// annouce applies the pull to all cards in the round.
func (r *Round) annouce(pull int) {
	for _, card := range r.cards {
		card.cover(pull)
	}
}

// winners returns the winning cards in the round.
func (r *Round) winners() []*Card {
	winners := make([]*Card, 0, len(r.cards))
	for _, card := range r.cards {
		if r.pattern.isWinner(card) {
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
