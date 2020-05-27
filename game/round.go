package game

// Round represents a round of BINGO, including the cage of numbers
// and any cards in play.
type Round struct {
	cage  *Cage
	cards []Card
}
