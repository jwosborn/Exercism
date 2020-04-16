package scrabble

import "strings"

var m = map[string]int{
	"A": 1,
	"B": 3,
	"C": 3,
	"D": 2,
	"E": 1,
	"F": 4,
	"I": 1,
	"O": 1,
	"U": 1,
	"L": 1,
	"N": 1,
	"R": 1,
	"S": 1,
	"T": 1,
	"G": 2,
	"M": 3,
	"P": 3,
	"V": 4,
	"W": 4,
	"Y": 4,
	"H": 4,
	"K": 5,
	"J": 8,
	"X": 8,
	"Q": 10,
	"Z": 10,
}

// Score iterates over a string and assigns a score to the string based on scrabble scoring
func Score(a string) int {
	score := 0
	for _, r := range a {
		letter := string(r)
		score += m[strings.ToUpper(letter)]
	}
	return score
}
