package game

import (
	"math/rand"
)

const count = 75
const rows = 5
const columns = 5
const free = 0
const nan = -1

type Round struct {
	cage  *Cage
	cards []Card
}

// Cage holds the Bingo balls numbered [1,75]
type Cage struct {
	Inside  []int
	Outside []int
}

// Cage co
type Card struct {
	B [rows]int // 1-15
	I [rows]int // 16-30
	N [rows]int // 31-45
	G [rows]int // 46-60
	O [rows]int // 61-75
}

func (c *Cage) IsEmpty() bool {
	return len(c.Inside) < 1
}

// Take returns a remaining number.
func (c *Cage) Take() int {
	if c.IsEmpty() {
		return nan
	}
	c.shuffle()                           // variety is the spice of life
	value := c.Inside[len(c.Inside)-1]    // take the last ball
	c.Inside = c.Inside[:len(c.Inside)-1] // remove the last ball Inside
	c.Outside = append(c.Outside, value)  // add the last ball Outside
	return value
}

func (s *Cage) shuffle() {
	rand.Shuffle(len(s.Inside), func(i, j int) {
		s.Inside[i], s.Inside[j] = s.Inside[j], s.Inside[i]
	})
}

// NewCage returns a new Bingo game
func NewCage() *Cage {
	var cage = &Cage{
		Inside: []int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
			16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
			46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
			61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75,
		},
		Outside: make([]int, 75),
	}
	cage.shuffle()
	return cage
}

func generateCardColumn(cage *Cage) [rows]int {
	var result [rows]int
	for i := 0; i < rows; i++ {
		result[i] = cage.Take()
	}
	return result
}

func (card *Card) fill() {
	cage := NewCage()
	card.B = generateCardColumn(cage)
	card.I = generateCardColumn(cage)
	card.N = generateCardColumn(cage)
	card.G = generateCardColumn(cage)
	card.O = generateCardColumn(cage)
}

// NewCard returns a Card
func NewCard() *Card {
	card := Card{}
	card.fill()
	return &card
}
