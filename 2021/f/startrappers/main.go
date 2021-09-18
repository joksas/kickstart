package startrappers

import (
	"fmt"
	"math"
)

type Coordinates struct {
	X int
	Y int
}

type ClosestPoint struct {
	Coordinates *Coordinates
	Distance    float64
}

func main() {
	var numCases int
	fmt.Scanln(&numCases)

	for idxCase := 1; idxCase <= numCases; idxCase++ {
		var numStars int
		fmt.Scanln(&numStars)
		var whiteStars = make([]Coordinates, numStars)
		for idxStar := 0; idxStar < numStars; idxStar++ {
			var x, y int
			fmt.Scan(&x, &y)
			whiteStar := Coordinates{X: x, Y: y}
			whiteStars[idxStar] = whiteStar
		}
		var x, y int
		fmt.Scan(&x, &y)
		blueStar := Coordinates{X: x, Y: y}
		if len(whiteStars) < 3 {
			printCase(idxCase, nil)
			continue
		}
		closestStars := FourClosest(blueStar, whiteStars)
		if len(closestStars) < 3 {
			printCase(idxCase, nil)
			continue
		}

		if len(closestStars) == 3 {
			for i := 0; i < 4; i++ {
				_, ok := closestStars[i]
				// Found missing star
				if !ok {
					// Found NE or SW
					if i%2 == 0 {
						if closestStars[1].Y == closestStars[3].Y {
							printCase(idxCase, nil)
							break
						}
						continue
					} else { // Found NW or SE
						if closestStars[0].X == closestStars[2].X {
							printCase(idxCase, nil)
							break
						}
						continue
					}
				}
				distance := StarsPerimeter(closestStars)
				printCase(idxCase, &distance)
				break
			}
			continue
		}

		distance := StarsPerimeter(closestStars)
		printCase(idxCase, &distance)
	}
}

func FourClosest(blueStar Coordinates, whiteStars []Coordinates) map[int]Coordinates {
	var NE = ClosestPoint{}
	var NW = ClosestPoint{}
	var SW = ClosestPoint{}
	var SE = ClosestPoint{}

	for _, whiteStar := range whiteStars {
		distance := distanceBetween(whiteStar, blueStar)
		// Each quadrant will take a border in counter-clockwise direction.
		if whiteStar.X >= blueStar.X && whiteStar.Y > blueStar.Y {
			NE = updateClosestPoint(whiteStar, distance, NE)
		}
		if whiteStar.X < blueStar.X && whiteStar.Y >= blueStar.Y {
			NW = updateClosestPoint(whiteStar, distance, NW)
		}
		if whiteStar.X <= blueStar.X && whiteStar.Y < blueStar.Y {
			SW = updateClosestPoint(whiteStar, distance, SW)
		}
		if whiteStar.X > blueStar.X && whiteStar.Y <= blueStar.Y {
			SE = updateClosestPoint(whiteStar, distance, SE)
		}
	}

	fourClosest := map[int]Coordinates{}

	if NE.Coordinates != nil {
		fourClosest[0] = *NE.Coordinates
	}
	if NW.Coordinates != nil {
		fourClosest[1] = *NW.Coordinates
	}
	if SW.Coordinates != nil {
		fourClosest[2] = *SW.Coordinates
	}
	if SE.Coordinates != nil {
		fourClosest[3] = *SE.Coordinates
	}

	return fourClosest
}

func updateClosestPoint(blueStar Coordinates, distance float64, closestPoint ClosestPoint) ClosestPoint {
	if closestPoint.Coordinates == nil || distance < closestPoint.Distance {
		closestPoint = ClosestPoint{
			Coordinates: &blueStar,
			Distance:    distance,
		}
	}

	return closestPoint
}

func distanceBetween(p1, p2 Coordinates) float64 {
	first := math.Pow(float64(p2.X-p1.X), 2)
	second := math.Pow(float64(p2.Y-p1.Y), 2)
	return math.Sqrt(first + second)
}

func printCase(idxCase int, distance *float64) {
	if distance == nil {
		fmt.Printf("Case #%d: IMPOSSIBLE\n", idxCase)
		return
	}
	fmt.Printf("Case #%d: %f\n", idxCase, *distance)
}

func StarsPerimeter(stars map[int]Coordinates) float64 {
	var perimeter float64

	var firstStar, lastStar Coordinates

	var startIdx = 1
	previousStar, ok := stars[0]
	if !ok {
		previousStar, _ = stars[1]
		startIdx = 2
	}
	firstStar = previousStar

	for i := startIdx; i < 4; i++ {
		currentStar, ok := stars[i]
		if !ok {
			continue
		}
		distance := distanceBetween(previousStar, currentStar)
		perimeter += distance
		previousStar = currentStar
		lastStar = currentStar
	}

	perimeter += distanceBetween(firstStar, lastStar)

	return perimeter
}
