package game

import (
	"errors"
	"math/rand"
)

const nan = -1         // invalid number
const standardMin = 1  // standard minimum value in BINGO
const standardMax = 75 // standard maximum value in BINGO

// Cage holds the BINGO balls for a round.
type Cage struct {
	Inside  []int
	Outside []int
}

// IsEmpty returns true if the cage has no remaining numbers inside.
func (c *Cage) IsEmpty() bool {
	return len(c.Inside) < 1
}

// Take returns and moves a number from inside to outside the cage.
func (c *Cage) Take() (int, error) {
	if c.IsEmpty() {
		return nan, errors.New("Cannot take from empty cage")
	}
	c.shuffle()                           // variety is the spice of life
	value := c.Inside[len(c.Inside)-1]    // take the last ball
	c.Inside = c.Inside[:len(c.Inside)-1] // remove the last ball Inside
	c.Outside = append(c.Outside, value)  // add the last ball Outside
	return value, nil
}

// shuffle randomizes the numbers inside the cage.
func (s *Cage) shuffle() {
	rand.Shuffle(len(s.Inside), func(i, j int) {
		s.Inside[i], s.Inside[j] = s.Inside[j], s.Inside[i]
	})
}

// NewCage returns a new cage with all the numbers in the range [min,max].
func NewCage(min, max int) *Cage {
	inside := make([]int, 0, max)
	for val := min; val <= max; val++ {
		inside = append(inside, val)
	}
	var cage = &Cage{
		Inside:  inside,
		Outside: make([]int, 0, max),
	}
	cage.shuffle()
	return cage
}

// NewStandardCage returns a new cage with numbers in the standard range [1,75].
func NewStandardCage() *Cage {
	return NewCage(standardMin, standardMax)
}
