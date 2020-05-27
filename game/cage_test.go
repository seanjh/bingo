package game

import "testing"

const ExpectedCageTotal = 2850 // 1 + 2 + 3 + ... + 75

func TestNewStandardCage(t *testing.T) {
	cage := NewStandardCage()
	if len(cage.Inside) != 75 {
		t.Errorf("Cage len = %d; want 75", len(cage.Inside))
	}

	sum := 0
	for _, val := range cage.Inside {
		sum += val
	}
	if sum != ExpectedCageTotal {
		t.Errorf("Cage total = %d; want %d", sum, ExpectedCageTotal)
	}
}
