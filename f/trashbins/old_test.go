package trashbins_test

import (
	"go/f/trashbins"
	"testing"
)

func TestTotalDistance(t *testing.T) {
	var tests = []struct {
		hasBinSlice   []bool
		totalDistance int
	}{
		{[]bool{true, true, true}, 0},
		{[]bool{true, false, false, true, false, false}, 5},
		{[]bool{false, false, true}, 3},
		{[]bool{false, true, false}, 2},
		{[]bool{true}, 0},
	}

	for i, test := range tests {
		totalDistance := trashbins.TotalDistance(test.hasBinSlice)
		if totalDistance != test.totalDistance {
			t.Errorf("test #%d: input=%v, expected=%d, got=%d", i, test.hasBinSlice, test.totalDistance, totalDistance)
		}
	}
}
