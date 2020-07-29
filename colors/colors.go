package colors

import (
	"math/rand"
	"time"
)

var colors = []string{
	"white",
	"black",
	"yellow",
	"green",
	"red",
	"blue",
	"gray",
	"brown",
}

// 0 - no match
// 1 - match
// 2 - match and same position
func CheckColors(yourColors []int, setColors[]int) []int {
	// instantiate the array with no matching elements
	result := []int{0,0,0,0,0}

	// checks if there are any matching colors
	for iYour, yourColor := range yourColors {
		for iSet, setColor := range setColors {
			if yourColor == setColor {
				// if match found - assign 1
				result[iYour] = 1
				// if match of color and position found - assign 2
				if iYour == iSet && yourColors[iYour] == setColors[iSet] {
					result[iYour] = 2
				}
			}
		}
	}

	return result
}

// SetColors sets the game by initializing five initial random colors
// every color must be different
func SetColors() []string {
	var usedNumbers []int
	rand.Seed(time.Now().UnixNano())
	for len(usedNumbers) < 5 {
		randomNumber := rand.Intn(8)
		if !checkNumber(usedNumbers, randomNumber) {
			usedNumbers = append(usedNumbers, randomNumber)
		}
	}

	var setColors []string
	for _, setColorAsNum := range usedNumbers {
		setColors = append(setColors, colors[setColorAsNum])
	}

	return setColors
}

func checkNumber(haystack []int, needle int) bool {
	for _, num := range haystack {
		if num == needle {
			return true
		}
	}

	return false
}

