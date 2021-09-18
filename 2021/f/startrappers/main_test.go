package startrappers_test

import (
	"go/2021/f/startrappers"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestFourClosest(t *testing.T) {
	var tests = []struct {
		blueStar    startrappers.Coordinates
		whiteStars  []startrappers.Coordinates
		fourClosest map[int]startrappers.Coordinates
	}{
		{
			startrappers.Coordinates{1, 1},
			[]startrappers.Coordinates{
				{0, 0},
				{5, 0},
				{0, 5},
			},
			map[int]startrappers.Coordinates{
				2: {0, 0},
				3: {5, 0},
				1: {0, 5},
			},
		},
		{
			startrappers.Coordinates{1, 1},
			[]startrappers.Coordinates{
				{0, 0},
				{-10, -10},
				{10, 0},
				{5, 0},
				{0, 5},
				{-5, 12},
			},
			map[int]startrappers.Coordinates{
				2: {0, 0},
				3: {5, 0},
				1: {0, 5},
			},
		},
	}

	for i, test := range tests {
		fourClosest := startrappers.FourClosest(test.blueStar, test.whiteStars)
		if !cmp.Equal(fourClosest, test.fourClosest) {
			t.Errorf("test #%d: input=(%+v, %+v), expected=%+v, got=%+v", i, test.blueStar, test.whiteStars, test.fourClosest, fourClosest)
		}
	}
}

func TestStarsPerimeter(t *testing.T) {
	var tests = []struct {
		stars     map[int]startrappers.Coordinates
		perimeter float64
	}{
		{
			map[int]startrappers.Coordinates{
				2: {0, 0},
				3: {5, 0},
				1: {0, 5},
			},
			17.071068,
		},
	}

	for i, test := range tests {
		perimeter := startrappers.StarsPerimeter(test.stars)
		if !cmp.Equal(perimeter, test.perimeter, cmpopts.EquateApprox(0.0001, 0.0001)) {
			t.Errorf("test #%d: input=%+v, expected=%f, got=%f", i, test.stars, test.perimeter, perimeter)
		}
	}
}
