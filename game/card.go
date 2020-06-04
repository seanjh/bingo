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

// Return the column label (e.g. "B") for the column number (1-B, 2-I, 3-N, 4-G, 5-O).
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

// parseCellName returns the column label and row number for a given cell name (e.g., "B1" -> ("B", 1))
func parseCellName(cellName string) (string, int, error) {
	var (
		colName string
		row     int
	)
	_, err := fmt.Sscanf(cellName, "%1s%d", &colName, &row)
	if err != nil {
		return "", nan, fmt.Errorf("failed to parse cell %s: %v", cellName, err)
	}
	return colName, row, nil
}

// cell is one number on a BINGO card.
type cell struct {
	column  int
	value   int
	covered bool
}

func (c *cell) String() string {
	column := getColumnLabel(c.column)
	state := ""
	if c.covered {
		state = " - X"
	}
	return fmt.Sprintf("%s%d%s", column, c.value, state)
}

// Cover sets this cell to covered.
func (c *cell) Cover() {
	c.covered = true
}

// column is one column of numbers on a single BINGO card.
type column struct {
	number int
	values []cell
}

func (c *column) label() string {
	return getColumnLabel(c.number)
}

func (c *column) getLastRowNum() int {
	return len(c.values) + 1
}

func (c *column) isValidRowNum(rowNum int) bool {
	return rowNum <= c.getLastRowNum()
}

func (c *column) cellAt(rowNum int) (cell, error) {
	if !c.isValidRowNum(rowNum) {
		lastLabel := fmt.Sprintf("%s%d", c.label(), c.getLastRowNum())
		badLabel := fmt.Sprintf("%s%d", c.label(), rowNum)
		return cell{}, fmt.Errorf("cannot access row %s beyond last row: %s", badLabel, lastLabel)
	}
	return c.values[rowNum-1], nil
}

// validRange returns the column's valid range as [lower,upper].
//
// For example, a 5x5 card with 3x multiple:
//	B (1) returns	1,15
//	I (2) returns	16,30
//	...
func (c *column) validRange(rows, multiple int) (int, int) {
	lower := rows*multiple*(c.number-1) + 1
	upper := rows * multiple * c.number
	return lower, upper
}

// fill populates the numbers in a column from the valid range.
func (c *column) fill(rows, multiple int) []cell {
	c.values = make([]cell, 0, rows)
	cage := NewCage(c.validRange(rows, multiple))
	for i := 0; i < rows; i++ {
		value, _ := cage.Take() // we're careful to avoid empty cages
		l := cell{column: c.number, value: value}
		c.values = append(c.values, l)
	}
	return c.values
}

func (c *column) addFreeSlot() {
	fmt.Println("TODO")
}

func (c *column) String() string {
	values := make([]string, 0, len(c.values))
	for _, cell := range c.values {
		values = append(values, cell.String())
	}
	return fmt.Sprintf("%s: %s", c.label(), strings.Join(values, ", "))
}

// Card contains 5 columns of randomized values.
type Card struct {
	B        *column
	I        *column
	N        *column
	G        *column
	O        *column
	rows     int
	multiple int
}

// newColumn returns the randomized column for the given column number.
func (card *Card) newColumn(colNum int) *column {
	col := &column{number: colNum}
	col.fill(card.rows, card.multiple)
	if colNum == N {
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

func (card *Card) columnFrom(colLabel string) (*column, error) {
	switch colLabel {
	case "B":
		return card.B, nil
	case "I":
		return card.I, nil
	case "N":
		return card.N, nil
	case "G":
		return card.G, nil
	case "O":
		return card.O, nil
	}
	return &column{}, fmt.Errorf("column label %s is unrecognized", colLabel)
}

// cellAt returns the cell referenced by the cell name (e.g, "B1").
func (card *Card) cellAt(cellName string) (cell, error) {
	colName, rowNum, err := parseCellName(cellName)
	if err != nil {
		return cell{}, err
	}

	if rowNum > card.rows {
		return cell{}, fmt.Errorf("parsed row %d > card rows %d", rowNum, card.rows)
	}

	col, err := card.columnFrom(colName)
	if err != nil {
		return cell{}, err
	}

	result, err := col.cellAt(rowNum)
	if err != nil {
		return cell{}, err
	}

	return result, nil
}

// ValueAt returns the value in the specified cell (e.g, "B1").
func (card *Card) ValueAt(cellName string) (int, error) {
	cell, err := card.cellAt(cellName)
	if err != nil {
		return nan, err
	}
	return cell.value, nil
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
