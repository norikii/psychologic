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

// 0 - not in array
// 1 - present in array
func CheckColors(yourColors []int, setColors[]int) []int {
	var result []int

	for _, yourColor := range yourColors {
		for _, setColor := range setColors {
			if yourColor == setColor {
				result = append(result, 1)
			}
		}
	}

	return result
}

func checkAll(yourColors []int, setColors[]int) []int {
	var result []int
	for i:=0; i<5; i++ {
		if yourColors[i] == setColors[i] {
			result = append(result, yourColors[i])
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

