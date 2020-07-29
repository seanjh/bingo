package game

import (
	"fmt"
)

const numColumns = 5          // fixed number of columns
const standardNumRows = 5     // default number of rows per column
const standardMultiple = 3    // default multiple of available to possible row values
const free = 0                // card "free" space
const invalidColumnName = "_" // invalid column name
const invalidColumnNum = -99

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

// addFreeSlot sets the center "N" cell to free.
func addFreeSlot(col []*cell) {
	l := freeCell(col)
	l.markFree()
}

// Card contains 5 columns of randomized values.
type Card struct {
	rows    [][]*cell
	columns [][]*cell
	//lookup   map[string]*cell  (e.g., B1)
}

// B returns the slice containing the B column's cells.
func (card *Card) B() []*cell { return card.columns[B-1] }

// I returns the slice containing the I column's cells.
func (card *Card) I() []*cell { return card.columns[I-1] }

// N returns the slice containing the N column's cells.
func (card *Card) N() []*cell { return card.columns[N-1] }

// G returns the slice containing the G column's cells.
func (card *Card) G() []*cell { return card.columns[G-1] }

// O returns the slice containing the O column's cells.
func (card *Card) O() []*cell { return card.columns[O-1] }

// lastRowNum returns the last row number on the card.
func (card *Card) lastRowNum() int {
	return len(card.rows)
}

// isValidRowNum returns true if the row number exists in the card
func (card *Card) isValidRow(rowNum int) bool {
	return 0 <= rowNum && rowNum <= card.lastRowNum()
}

// getRow returns the specified row from the card (if present)
func (card *Card) getRow(rowNum int) (error, []*cell) {
	if card.isValidRow(rowNum) {
		return nil, card.rows[rowNum]
	}
	return fmt.Errorf("Invalid row number: %d from %d", rowNum, len(card.rows)), []*cell{}
}

// fillCells populates the cells with values from a new cage.
func (card *Card) fillCol(cells []*cell, cage *Cage, colNum int) {
	for i := 0; i < len(cells); i++ {
		value, _ := cage.Take() // we're careful to avoid empty cages
		l := &cell{
			column: colNum,
			row:    i + 1,
			value:  value,
		}
		cells[i] = l
	}
}

func (card *Card) rowFromColumns(rowNum int) []*cell {
	return []*cell{
		card.B()[rowNum-1],
		card.I()[rowNum-1],
		card.N()[rowNum-1],
		card.G()[rowNum-1],
		card.O()[rowNum-1],
	}
}

// fill populates the BINGO card columns with values in their valid range.
func (card *Card) fill(numRows, multiple int) {
	// TODO(sean): tidy
	card.columns = make([][]*cell, numColumns)
	for colNum := B; colNum <= O; colNum++ {
		// create new column
		col := make([]*cell, numRows)
		cage := NewCage(validColumnRange(colNum, numRows, multiple))
		card.fillCol(col, cage, colNum)
		if colNum == N {
			addFreeSlot(col)
		}
		card.columns[colNum-1] = col
	}

	card.rows = make([][]*cell, numRows)
	for i := 0; i < numRows; i++ {
		// create row from columns
		card.rows[i] = card.rowFromColumns(i + 1)
	}
}

// cellAt returns the cell referenced by the cell name (e.g, "B1").
func (card *Card) cellAt(cellName string) (*cell, error) {
	colName, rowNum, err := parseCellName(cellName)
	if err != nil {
		return &cell{}, err
	}

	if last := card.lastRowNum(); rowNum > last {
		return &cell{}, fmt.Errorf("parsed row %d from cell '%s'; want row <= %d", rowNum, cellName, last)
	} else if rowNum < 1 {
		return &cell{}, fmt.Errorf("parsed row %d from cell '%s'; want > 1", rowNum, cellName)
	}

	colNum := getColumnNum(colName)
	if colNum == invalidColumnNum {
		return &cell{}, fmt.Errorf("Invalid column for cell '%s'", cellName)
	}

	return card.rows[rowNum-1][colNum-1], nil
}

// ValueAt returns the value in the specified cell (e.g, "B1").
func (card *Card) ValueAt(cellName string) (int, error) {
	cell, err := card.cellAt(cellName)
	if err != nil {
		return nan, err
	}
	return cell.value, nil
}

// IsWinner returns true if the card matches the pattern.
func (*Card) IsWinner() bool {
	fmt.Println("Not Implemeneted")
	return false
}

// NewCard returns a new Card with 5 columns (B, I, N, G, O), the specified
// number of rows, and values randomly populated from the range [1,5*rows*multiple].
func NewCard(numRows, multiple int) *Card {
	card := &Card{}
	card.fill(numRows, multiple)
	return card
}

// NullCard returns a new empty Card.
func NullCard() *Card {
	return &Card{[][]*cell{}, [][]*cell{}}
}

// NewStandardCard returns a standard 5x5 BINGO card.
func NewStandardCard() *Card {
	return NewCard(standardNumRows, standardMultiple)
}
