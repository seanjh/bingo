package game

import "testing"

const ExpectedCageTotal = 2850 // 1 + 2 + 3 + ... + 75

func TestNewStandardCage(t *testing.T) {
	cage := NewStandardCage()
	if len(cage.Inside) != 75 {
		t.Errorf("Cage len Inside = %d; want 75", len(cage.Inside))
	}

	sum := 0
	for _, val := range cage.Inside {
		sum += val
	}
	if sum != ExpectedCageTotal {
		t.Errorf("Cage total = %d; want %d", sum, ExpectedCageTotal)
	}
}

func TestCageEmpty(t *testing.T) {
	cage := NewCage(1, 1)
	if len(cage.Inside) != 1 {
		t.Errorf("Cage len Inside = %d; want 1", len(cage.Inside))
	}

	val, err := cage.Take()
	if err != nil {
		t.Errorf("Take() err = %v; want nil", err)
	}

	if val != 1 {
		t.Errorf("Take() value = %d; want 1", val)
	}

	if len(cage.Inside) != 0 {
		t.Errorf("Cage len Inside = %d; want 0", len(cage.Inside))
	}

	if len(cage.Outside) != 1 {
		t.Errorf("Cage len Outside = %d; want 1", len(cage.Outside))
	}

	if !cage.IsEmpty() {
		t.Error("Cage IsEmpty = false; want true")
	}
}

func TestTakeEmpty(t *testing.T) {
	cage := NewCage(1, 1)
	cage.Take()
	_, err := cage.Take()
	if err == nil {
		t.Error("Take() err is nil; want error for empty cage")
	}
}

func TestTakeSolitary(t *testing.T) {
	cases := []struct {
		cage     *Cage
		expected int
	}{
		{NewCage(1, 1), 1},
		{NewCage(5, 5), 5},
		{NewCage(0, 0), 0},
		//{NewCage(-1, -1), -1}, TODO: maybe fix explosion when < 0
	}

	for _, c := range cases {
		actual, err := c.cage.Take()
		if err != nil {
			t.Errorf("cage err = %v; want nil", err)
		}
		if actual != c.expected {
			t.Errorf("cage Take() = %d; want %d", actual, c.expected)
		}
	}
}

func TestTakeAll(t *testing.T) {
	cases := []struct {
		cage     *Cage
		possible []int
	}{
		{NewCage(1, 2), []int{1, 2}},
		{NewCage(1, 5), []int{1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		actual := make([]int, 0, len(c.possible))

		// take all values
		for !c.cage.IsEmpty() {
			value, _ := c.cage.Take()
			actual = append(actual, value)
		}

		// ensure all values taken from the cage are valid
		for _, val := range actual {
			found := false
			for _, expected := range c.possible {
				if val == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("cage %v missing %d", c.possible, val)
			}
		}
	}
}
