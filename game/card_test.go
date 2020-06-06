package game

import (
	"fmt"
	"testing"
)

func TestGetColumnLabel(t *testing.T) {
	cases := []struct {
		number   int
		expected string
	}{
		{B, "B"},
		{I, "I"},
		{N, "N"},
		{G, "G"},
		{O, "O"},
		{0, ""},
		{-1, ""},
		{99, ""},
	}

	for _, c := range cases {
		if actual := getColumnLabel(c.number); actual != c.expected {
			t.Errorf("getColumnLabel(%d) = %s; want %s", c.number, actual, c.expected)
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

func TestColumnLabel(t *testing.T) {
	cases := []struct {
		col      column
		expected string
	}{
		{column{B, []cell{}}, "B"},
		{column{I, []cell{}}, "I"},
		{column{N, []cell{}}, "N"},
		{column{G, []cell{}}, "G"},
		{column{O, []cell{}}, "O"},
		{column{0, []cell{}}, ""},
		{column{99, []cell{}}, ""},
		{column{-1, []cell{}}, ""},
	}

	for _, c := range cases {
		if actual := c.col.label(); actual != c.expected {
			t.Errorf("col.label() = %s; want %s", actual, c.expected)
		}
	}
}

func TestStandardCard(t *testing.T) {
	card := NewStandardCard()

	if card.rows != 5 {
		t.Errorf("standard card rows = %d; want 5", card.rows)
	}
	if card.multiple != 3 {
		t.Errorf("standard card multiple = %d; want 3", card.multiple)
	}

	cases := []struct {
		colName  string
		values   []cell
		expected int
	}{
		{"B", card.B.values, 5},
		{"I", card.I.values, 5},
		{"N", card.N.values, 5},
		{"G", card.G.values, 5},
		{"O", card.O.values, 5},
	}

	for _, c := range cases {
		if actual := len(c.values); actual != c.expected {
			t.Errorf("len(%s) = %d; want %d", c.colName, actual, c.expected)
		}
	}
}

func TestCardValueAt(t *testing.T) {
	cases := []struct {
		cellName string
		expected int
	}{
		{"B1", 1}, {"I1", 2}, {"N1", free}, {"G1", 4}, {"O1", 5},
		//{"B2", 2}, {"I2", 2}, {"N2", free}, {"G2", 2}, {"O2", 2},
		//{"B3", 3}, {"I3", 3}, {"N3", 3}, {"G3", 3}, {"O3", 3},
	}

	card := NewCard(1, 1)
	for _, c := range cases {
		if actual, err := card.ValueAt(c.cellName); err != nil {
			t.Errorf("ValueAt(%s) err: '%v'", c.cellName, err)
		} else if actual != c.expected {
			t.Errorf("%s is %d; want %d", c.cellName, actual, c.expected)
		}
	}
}

func TestCardValueAtError(t *testing.T) {
	cases := []string{
		"B99",
		"B0",
		"",
		"Z1",
		"1B",
		"FOO",
	}

	card := NewCard(1, 1)
	for _, cellName := range cases {
		actual, err := card.ValueAt(cellName)
		if err == nil {
			t.Errorf("card.ValueAt(%s) = %d; want err", cellName, actual)
		}
	}
}

func TestFreeCell(t *testing.T) {
	cases := []struct {
		col      column
		expected *cell
	}{}

	for _, c := range cases {
		if actual := c.col.freeCell(); actual != c.expected {
			t.Errorf("col.freeCell() = %v; want %v", actual, c.expected)
		}
	}
}

func TestCellCovered(t *testing.T) {
	cases := []struct {
		cell     *cell
		expected string
	}{
		{&cell{}, "0"},
		{&cell{value: 1}, "1"},
		{&cell{column: B, value: 1}, "B1"},
		{&cell{column: B, value: 1, covered: true}, "B1 - X"},
	}

	for _, c := range cases {
		actual := fmt.Sprintf("%s", c.cell)
		if actual != c.expected {
			t.Errorf("cell.String() = %s; want %s", actual, c.expected)
		}
	}
}
