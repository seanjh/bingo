package game

import (
	"fmt"
	"testing"
)

func TestCellCovered(t *testing.T) {
	cases := []struct {
		cell     *cell
		expected string
	}{
		{&cell{}, "_0:0"},
		{&cell{value: 1}, "_0:1"},
		{&cell{column: B, row: 1, value: 1}, "B1:1"},
		{&cell{column: B, row: 1, value: 1, covered: true}, "B1|1"},
		{&cell{column: O, row: 20, value: 50}, "O20:50"},
	}

	for _, c := range cases {
		actual := fmt.Sprintf("%s", c.cell)
		if actual != c.expected {
			t.Errorf("cell.String() = %s; want %s", actual, c.expected)
		}
	}
}

func TestParseCellName(t *testing.T) {
	cases := []struct {
		cellName string
		colLabel string
		rowNum   int
	}{
		{"B1", "B", 1},
		{"B0", "B", 0},
		{"A99", "A", 99},
		{"Z-1", "Z", -1},
	}

	for _, c := range cases {
		colLabel, rowNum, err := parseCellName(c.cellName)
		if err != nil {
			t.Errorf("parseCellName err = %v; want nil", err)
		}
		if colLabel != c.colLabel {
			t.Errorf("parseCellName colLabel = %s; want %s", colLabel, c.colLabel)
		}
		if rowNum != c.rowNum {
			t.Errorf("parseCellName rowNum = %d; want %d", rowNum, c.rowNum)
		}
	}
}

func TestParseCellNameErrors(t *testing.T) {
	cases := []string{
		"X",
		"",
		"ABC",
	}

	for _, cellName := range cases {
		colName, rowNum, err := parseCellName(cellName)
		if err == nil {
			t.Errorf(
				"parseCellName err = nil, colName = %s, rowNum = %d; want err",
				colName, rowNum)
		}
	}
}
