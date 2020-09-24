package game

import (
	"testing"
)

func TestGetColumnName(t *testing.T) {
	cases := []struct {
		number   int
		expected string
	}{
		{B, "B"},
		{I, "I"},
		{N, "N"},
		{G, "G"},
		{O, "O"},
		{0, "_"},
		{-1, "_"},
		{99, "_"},
	}

	for _, c := range cases {
		if actual := getColumnName(c.number); actual != c.expected {
			t.Errorf("getColumnName(%d) = %s; want %s", c.number, actual, c.expected)
		}
	}
}

func TestCellAt(t *testing.T) {
	cases := []struct {
		cellName string
		expected cell
	}{
		{"", cell{}},
	}

	card := &Card{}
	for _, c := range cases {
		if actual, _ := card.cellAt(c.cellName); *actual != c.expected {
			t.Errorf("card cell at '%s' = %v; want %v", c.cellName, actual, c.expected)
		}
	}
}

func TestStandardCard(t *testing.T) {
	card := NewStandardCard()

	if last := card.lastRowNum(); last != 5 {
		t.Errorf("standard card last row = %d; want 5", last)
	}

	cases := []struct {
		colName  string
		values   []*cell
		expected int
	}{
		{"B", card.B(), 5},
		{"I", card.I(), 5},
		{"N", card.N(), 5},
		{"G", card.G(), 5},
		{"O", card.O(), 5},
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
		col      []*cell
		expected *cell
	}{}

	for _, c := range cases {
		if actual := freeCell(c.col); actual != c.expected {
			t.Errorf("freeCell = %v; want %v", actual, c.expected)
		}
	}
}
