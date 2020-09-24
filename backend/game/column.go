package game

import (
	"math"
)

const (
	// B column on a BINGO card
	B = 1
	// I column on a BINGO card
	I = 2
	// N column on a BINGO card
	N = 3
	// G column on a BINGO card
	G = 4
	// O column on a BINGO card
	O = 5
)

// addFreeSlot sets the center "N" cell to free.
func addFreeSlot(col []*cell) {
	l := freeCell(col)
	l.markFree()
}

// Return the column label (e.g. "B") for the column number (1-B, 2-I, 3-N, 4-G, 5-O).
func getColumnName(number int) string {
	switch number {
	case B:
		return "B"
	case I:
		return "I"
	case N:
		return "N"
	case G:
		return "G"
	case O:
		return "O"
	}
	return "_"
}

// getColumnName returns the column number from the column name.
func getColumnNum(colName string) int {
	switch colName {
	case "B":
		return B
	case "I":
		return I
	case "N":
		return N
	case "G":
		return G
	case "O":
		return O
	}
	return invalidColumnNum
}

// freeCell returns the center cell in the column (shifted left for even len columns).
func freeCell(col []*cell) *cell {
	n := len(col)
	if n == 0 {
		return &cell{}
	}

	var i int
	if n%2 == 0 {
		i = n/2 - 1 // even
	} else {
		i = n / 2 // odd
	}

	return col[i] // odd
}

// validColumnRange returns the column's valid range as [lower,upper].
//
// For example, a 5x5 card with 3x multiple:
//	B (1) returns	1,15
//	I (2) returns	16,30
//	...
func validColumnRange(colNum, numRows, multiple int) (int, int) {
	lower := numRows*multiple*(colNum-1) + 1
	upper := numRows * multiple * colNum
	return lower, upper
}

func columnForPull(pull int, card *Card) int {
	raw := float64(pull) / float64(len(card.rows)) / float64(card.multiple)
	colNum := math.Ceil(raw)
	return int(colNum)
}
