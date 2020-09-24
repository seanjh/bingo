package game

import (
	"fmt"
)

// cell is one square on a BINGO card.
type cell struct {
	column  int
	row     int
	value   int
	covered bool
}

// String returns a string representation of a cell.
func (c *cell) String() string {
	state := ":"
	if c.covered {
		state = "|"
	}
	return fmt.Sprintf("%s%d%s%d", getColumnName(c.column), c.row, state, c.value)
}

// Cover sets this cell to covered.
func (c *cell) Cover() {
	c.covered = true
}

// markFree sets cell value to free.
func (c *cell) markFree() {
	c.value = free
}

// parseCellName returns the column label and row number for a given cell name
// (e.g., "B1" -> ("B", 1))
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
