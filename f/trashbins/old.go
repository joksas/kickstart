package trashbins

import "sort"

func hasBinStrToBoolSlice(hasBinStr string) []bool {
	hasBinSlice := make([]bool, len(hasBinStr))

	for idx, char := range hasBinStr {
		if char == '0' {
			hasBinSlice[idx] = false
		} else {
			hasBinSlice[idx] = true
		}
	}

	return hasBinSlice
}

func TotalDistance(hasBinSlice []bool) int {
	var binPositions []int
	for idx, hasBin := range hasBinSlice {
		if hasBin {
			binPositions = append(binPositions, idx)
		}
	}

	var totalDistance int
	for pos, hasBin := range hasBinSlice {
		if hasBin {
			continue
		}
		idxRight := sort.SearchInts(binPositions, pos)
		if idxRight >= len(binPositions) {
			totalDistance += pos - binPositions[len(binPositions)-1]
			continue
		}
		if idxRight == 0 {
			totalDistance += binPositions[0] - pos
			continue
		}
		posRight := binPositions[idxRight]
		posLeft := binPositions[idxRight-1]

		distanceLeft := pos - posLeft
		distanceRight := posRight - pos
		if distanceLeft < distanceRight {
			totalDistance += distanceLeft
		} else {
			totalDistance += distanceRight
		}
	}

	return totalDistance
}
