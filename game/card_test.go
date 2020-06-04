package game

import (
	"fmt"
	"testing"
)

func TestStandardCard(t *testing.T) {
	card := NewStandardCard()

	if card.rows != 5 {
		t.Errorf("standard card rows = %d; want 5", card.rows)
	}
	if card.multiple != 3 {
		t.Errorf("standard card multiple = %d; want 3", card.multiple)
	}

	if len(card.B.values) != 5 {
		t.Errorf("len(card.B.values) = %d; want 5", len(card.B.values))
	}
	if len(card.I.values) != 5 {
		t.Errorf("len(card.I.values) = %d; want 5", len(card.I.values))
	}
	if len(card.N.values) != 5 {
		t.Errorf("len(card.N.values) = %d; want 5", len(card.N.values))
	}
	if len(card.G.values) != 5 {
		t.Errorf("len(card.G.values) = %d; want 5", len(card.G.values))
	}
	if len(card.O.values) != 5 {
		t.Errorf("len(card.O.values) = %d; want 5", len(card.O.values))
	}
}

func TestCardValueAt(t *testing.T) {
	cases := []struct {
		cellName string
		expected int
	}{
		{"B1", 1}, {"I1", 2}, {"N1", 3}, {"G1", 4}, {"O1", 5},
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

func TestCellCovered(t *testing.T) {
	cases := []struct {
		cell     *cell
		expected string
	}{
		{&cell{}, "WTF0"},
		{&cell{value: 1}, "WTF1"},
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
