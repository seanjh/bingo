package game

import (
	"fmt"
)

const standardRows = 5     // default number of rows per column
const standardMultiple = 3 // default multiple of available to possible row values
const free = 0             // card "free" space

const (
	B = 1
	I = 2
	N = 3
	G = 4
	O = 5
)

func getColumnLabel(number int) string {
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
	return "WTF"
}

// Cell is one number on a BINGO card.
type Cell struct {
	column  int
	value   int
	covered bool
}

func (c *Cell) String() string {
	column := getColumnLabel(c.column)
	state := ""
	if c.covered {
		state = " - X"
	}
	return fmt.Sprintf("%s%d%s", column, c.value, state)
}

func (c *Cell) Cover() {
	c.covered = true
}

// Column is one column of numbers on a single BINGO card.
type Column struct {
	number int
	values []*Cell
}

// validRange returns the column's valid range as [lower,upper].
//
// For example, a 5x5 card with 3x multiple:
//	B (1) returns	1,15
//	I (2) returns	16,30
//	...
func (col *Column) validRange(rows, multiple int) (int, int) {
	lower := rows*multiple*(col.number-1) + 1
	upper := rows * multiple * col.number
	return lower, upper
}

// fill populates the numbers in a column from the valid range.
func (col *Column) fill(rows, multiple int) {
	col.values = make([]*Cell, 0, rows)

	cage := NewCage(col.validRange(rows, multiple))
	for i := 0; i < rows; i++ {
		value, _ := cage.Take() // we're careful to avoid empty cages
		cell := &Cell{column: col.number, value: value}
		col.values = append(col.values, cell)
	}
}

func (col *Column) addFreeSlot() {
	fmt.Println("TODO")
}

// Card contains 5 columns of randomized values.
type Card struct {
	B        *Column
	I        *Column
	N        *Column
	G        *Column
	O        *Column
	rows     int
	multiple int
}

// newColumn returns the randomized column for the given column number.
func (card *Card) newColumn(number int) *Column {
	col := &Column{number: number}
	col.fill(card.rows, card.multiple)

	if number == N {
		col.addFreeSlot()
	}

	return col
}

// fill populates the BINGO card columns with values in their valid range.
func (card *Card) fill() {
	card.B = card.newColumn(B)
	card.I = card.newColumn(I)
	card.N = card.newColumn(N)
	card.G = card.newColumn(G)
	card.O = card.newColumn(O)
}

// NewCard returns a new Card with 5 columns (B, I, N, G, O), the specified
// number of rows, and values randomly populated from the range [1,5*rows*multiple].
func NewCard(rows, multiple int) *Card {
	card := &Card{
		rows:     rows,
		multiple: multiple,
	}
	card.fill()
	return card
}

// NewStandardCard returns a standard 5x5 BINGO card.
func NewStandardCard() *Card {
	return NewCard(standardRows, standardMultiple)
}
