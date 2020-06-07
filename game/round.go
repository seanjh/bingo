package game

// Round represents a round of BINGO, including the cage of numbers
// and any cards in play.
type Round struct {
	cage  *Cage
	cards []*Card
}

// NewStandardRound returns a standard cage and cards for one round of BINGO.
func NewStandardRound(numCards int) *Round {
	cage := NewStandardCage()
	// TODO handle negative
	cards := make([]*Card, 0, numCards)
	for i := 0; i < numCards; i++ {
		cards = append(cards, NewStandardCard())
	}
	return &Round{cage: cage, cards: cards}
}
