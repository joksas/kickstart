package trashbins_test

import (
	"go/2021/f/trashbins"
	"testing"
)

func TestTotalDistanceEfficient(t *testing.T) {
	var tests = []struct {
		binPositions   []int
		noBinPositions []int
		totalDistance  int
	}{
		{[]int{0, 1, 2}, []int{}, 0},
		{[]int{0, 3}, []int{1, 2, 4, 5}, 5},
		{[]int{2}, []int{0, 1}, 3},
		{[]int{1}, []int{0, 2}, 2},
		{[]int{0}, []int{}, 0},
	}

	for i, test := range tests {
		totalDistance := trashbins.TotalDistanceEfficient(test.binPositions, test.noBinPositions)
		if totalDistance != test.totalDistance {
			t.Errorf("test #%d: input=(%v, %v), expected=%d, got=%d", i, test.binPositions, test.noBinPositions, test.totalDistance, totalDistance)
		}
	}
}
