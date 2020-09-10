package game

// any horizontal line (e.g., B1, I1, N1, G1, O1)
// any vertical line (e.g., B1, B2, B3, B4, B5)
// any diagonal line (e.g., B1, I2, free, G4, O5)
// blackout
//
// NEW HAMPSHIRE RACING AND CHARITABLE GAMING COMMISSIONAPPROVED GAME PATTERNS
// https://www.racing.nh.gov/forms-pubs/documents/game-patterns-approved.pdf

type Pattern interface {
	isWinner(*Card) bool
}

type any struct{}

func (p *any) isWinner(card *Card) bool { return true }

type horizontal struct{}

// isWinningRow returns true if all cells in the row are marked.
func (p *horizontal) isWinningRow(row []*cell) bool {
	for _, l := range row {
		if !l.covered {
			return false
		}
	}
	return true
}

// isWinner returns true and the winning cells if the card includes
// a row with all cells marked.
func (p *horizontal) isWinner(card *Card) bool {
	for _, row := range card.rows {
		if p.isWinningRow(row) {
			return true
		}
	}
	return false
}

// any line pattern
// composed of horizontal, vertical, and diagonal patterns
