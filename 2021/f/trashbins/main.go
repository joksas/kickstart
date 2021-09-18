package trashbins

import (
	"fmt"
	"sort"
)

func main() {
	var numCases int
	fmt.Scanln(&numCases)

	for idxCase := 1; idxCase <= numCases; idxCase++ {
		var dummy int
		fmt.Scanln(&dummy)
		var hasBinStr string
		fmt.Scanln(&hasBinStr)

		binPositions, noBinPositions := binAndNoBinPositions(hasBinStr)
		totalDistance := TotalDistanceEfficient(binPositions, noBinPositions)
		printCase(idxCase, totalDistance)
		if idxCase < numCases {
			fmt.Println()
		}
	}
}

func binAndNoBinPositions(hasBinStr string) ([]int, []int) {
	var binPositions []int
	var noBinPositions []int

	for idx, char := range hasBinStr {
		if char == '0' {
			noBinPositions = append(noBinPositions, idx)
		} else {
			binPositions = append(binPositions, idx)
		}
	}

	return binPositions, noBinPositions
}

func TotalDistanceEfficient(binPositions, noBinPositions []int) int {
	var totalDistance int

	for _, noBinPosition := range noBinPositions {
		idxRight := sort.SearchInts(binPositions, noBinPosition)
		if idxRight >= len(binPositions) {
			totalDistance += noBinPosition - binPositions[len(binPositions)-1]
			continue
		}
		if idxRight == 0 {
			totalDistance += binPositions[0] - noBinPosition
			continue
		}
		binPositionRight := binPositions[idxRight]
		binPositionLeft := binPositions[idxRight-1]

		distanceLeft := noBinPosition - binPositionLeft
		distanceRight := binPositionRight - noBinPosition
		if distanceLeft < distanceRight {
			totalDistance += distanceLeft
		} else {
			totalDistance += distanceRight
		}
	}

	return totalDistance
}

func printCase(idx, totalDistance int) {
	fmt.Printf("Case #%d: %d", idx, totalDistance)
}
