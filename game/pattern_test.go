package game

import "testing"

func TestLineAcross(t *testing.T) {
	cases := []int{0, 1, 2, 3, 4}
	for _, i := range cases {
		c := NewStandardCard()
		for _, l := range c.rows[i] {
			l.Cover()
		}
		isWinner := (&across{}).isWinner(c)
		if !isWinner {
			t.Errorf("card %v should have been a winner", c)
		}
	}
}
