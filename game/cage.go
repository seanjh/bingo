package game

import (
	"errors"
	"math/rand"
)

const nan = -1         // invalid number
const standardMin = 1  // standard minimum value in BINGO
const standardMax = 75 // standard maximum value in BINGO

var EmptyCage = errors.New("Empty cage")

// Cage holds the BINGO balls for a round.
type Cage struct {
	Inside  []int
	Outside []int
	rng     *rand.Rand
}

// NewCage returns a new cage with all the numbers in the range [min,max].
func NewCage(min, max int) *Cage {
	inside := make([]int, 0, max)
	for val := min; val <= max; val++ {
		inside = append(inside, val)
	}
	cage := &Cage{
		Inside:  inside,
		Outside: make([]int, 0, max),
		rng:     NewRng(),
	}
	cage.shuffle()
	return cage
}

// NewStandardCage returns a new cage with numbers in the standard range [1,75].
func NewStandardCage() *Cage {
	return NewCage(standardMin, standardMax)
}

// NullCage returns a new empty cage
func NullCage() *Cage {
	return &Cage{}
}

// IsEmpty returns true if the cage has no remaining numbers inside.
func (c *Cage) IsEmpty() bool {
	return len(c.Inside) < 1
}

// take returns and moves a number from inside to outside the cage.
func (c *Cage) take() (int, error) {
	if c.IsEmpty() {
		return nan, EmptyCage
	}
	c.shuffle()                           // variety is the spice of life
	value := c.Inside[len(c.Inside)-1]    // take the last ball
	c.Inside = c.Inside[:len(c.Inside)-1] // remove the last ball Inside
	c.Outside = append(c.Outside, value)  // add the last ball Outside
	return value, nil
}

// shuffle randomizes the numbers inside the cage.
func (c *Cage) shuffle() {
	c.rng.Shuffle(len(c.Inside), func(i, j int) {
		c.Inside[i], c.Inside[j] = c.Inside[j], c.Inside[i]
	})
}
