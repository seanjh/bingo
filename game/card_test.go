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

func TestCardFill(t *testing.T) {
	card := NewCard(1, 1)

	one := &cell{column: B, value: 1}
	// TODO(sean): card.B[0] would be a nicer way to access B1
	// this means card.B is []cell and Column goes away
	if actual := card.B.values[0]; actual.value != one.value {
		t.Errorf("B1 is '%s'; want '%s'", actual, one)
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
