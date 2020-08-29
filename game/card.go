package game

import (
	"fmt"
	"log"
)

const numColumns = 5          // fixed number of columns
const standardNumRows = 5     // default number of rows per column
const standardMultiple = 3    // default multiple of available to possible row values
const free = 0                // card "free" space
const invalidColumnName = "_" // invalid column name
const invalidColumnNum = -99

// Card contains 5 columns of randomized values.
type Card struct {
	rows    [][]*cell
	columns [][]*cell
	//lookup   map[string]*cell  (e.g., B1)
}

// NewCard returns a new Card with 5 columns (B, I, N, G, O), the specified
// number of rows, and values randomly populated from the range [1,5*rows*multiple].
func NewCard(numRows, multiple int) *Card {
	card := &Card{}
	card.fill(numRows, multiple)
	return card
}

// NewStandardCard returns a standard 5x5 BINGO card.
func NewStandardCard() *Card {
	return NewCard(standardNumRows, standardMultiple)
}

// NullCard returns a new empty Card.
func NullCard() *Card {
	return &Card{[][]*cell{}, [][]*cell{}}
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

// isWinner returns true if the card matches the pattern.
func (*Card) isWinner() bool {
	log.Print("Not Implemeneted")
	return true
}

// ValueAt returns the value in the specified cell (e.g, "B1").
func (card *Card) ValueAt(cellName string) (int, error) {
	cell, err := card.cellAt(cellName)
	if err != nil {
		return nan, err
	}
	return cell.value, nil
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

// TODO: Implement
func (c *Card) cover(pull int) {}

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

// fillCells populates the cells with values from a new cage.
func (card *Card) fillCol(cells []*cell, cage *Cage, colNum int) {
	for i := 0; i < len(cells); i++ {
		value, _ := cage.take() // we're careful to avoid empty cages
		l := &cell{
			column: colNum,
			row:    i + 1,
			value:  value,
		}
		cells[i] = l
	}
}

// getRow returns the specified row from the card (if present)
func (card *Card) getRow(rowNum int) (error, []*cell) {
	if card.isValidRow(rowNum) {
		return nil, card.rows[rowNum]
	}
	return fmt.Errorf("Invalid row number: %d from %d", rowNum, len(card.rows)), []*cell{}
}

// isValidRowNum returns true if the row number exists in the card
func (card *Card) isValidRow(rowNum int) bool {
	return 0 <= rowNum && rowNum <= card.lastRowNum()
}

// lastRowNum returns the last row number on the card.
func (card *Card) lastRowNum() int {
	return len(card.rows)
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
