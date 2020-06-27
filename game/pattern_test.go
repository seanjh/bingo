package game

import "testing"

func TestLineAcross(t *testing.T) {
	c := &Card{[][]*cell{}}
	isWinner := (&across{}).isWinner(c)
	if !isWinner {
		t.Errorf("card %v should have been a winner", c)
	}
}