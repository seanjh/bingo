package game

import (
	"fmt"
	"strings"
)

const standardRows = 5     // default number of rows per column
const standardMultiple = 3 // default multiple of available to possible row values
const free = 0             // card "free" space

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
	values []Cell
}

// validRange returns the column's valid range as [lower,upper].
//
// For example, a 5x5 card with 3x multiple:
//	B (1) returns	1,15
//	I (2) returns	16,30
//	...
func validRange(rows, multiple, colNum int) (int, int) {
	lower := rows*multiple*(colNum-1) + 1
	upper := rows * multiple * colNum
	return lower, upper
}

// fill populates the numbers in a column from the valid range.
func fillColumn(rows, multiple, colNum int) []Cell {
	col := make([]Cell, 0, rows)
	cage := NewCage(validRange(rows, multiple))
	for i := 0; i < rows; i++ {
		value, _ := cage.Take() // we're careful to avoid empty cages
		cell := &Cell{column: colNum, value: value}
		col = append(col, cell)
	}
	return col
}

func addFreeSlot(col []Cell) {
	fmt.Println("TODO")
}

type Column []Cell

// Card contains 5 columns of randomized values.
type Card struct {
	B        Column
	I        Column
	N        Column
	G        Column
	O        Column
	rows     int
	multiple int
}

// newColumn returns the randomized column for the given column number.
func (card *Card) newColumn(colNum int) *Column {
	col := fillColumn(card.rows, card.multiple, colNum)
	if number == N {
		addFreeSlot(col)
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

func parseCellName(card *Card, cell string) ([]Cell, int, error) {
	var (
		colName string
		row     int
	)
	_, err := fmt.Scanf(strings.NewReader(cell), "%s%d", &colName, &row)
	if err != nil {
		return make([]Cell), -1, errors.New("failed to parse cell: %s", cell)
	}
	return colName, row, nil
}

func (card *Card) ValueAt(cell string) (int, error) {
	col, row, err := parseCellName(card, cell)
	if err != nil {
		return -1, err
	}
	if row > card.rows {
		return -1, errors.New("parsed row %d > card rows %d", row, card.rows)
	}
	return col[row], nil
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
