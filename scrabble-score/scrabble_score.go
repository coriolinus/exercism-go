package scrabble

import (
	"strings"
)

// reverse scores (copy-paste-regex from README)
func reverseScores() map[int][]rune {
	return map[int][]rune{
		1:  []rune{rune('A'), rune('E'), rune('I'), rune('O'), rune('U'), rune('L'), rune('N'), rune('R'), rune('S'), rune('T')},
		2:  []rune{rune('D'), rune('G')},
		3:  []rune{rune('B'), rune('C'), rune('M'), rune('P')},
		4:  []rune{rune('F'), rune('H'), rune('V'), rune('W'), rune('Y')},
		5:  []rune{rune('K')},
		8:  []rune{rune('J'), rune('X')},
		10: []rune{rune('Q'), rune('Z')},
	}
}

func scores() map[rune]int {
	out := make(map[rune]int)
	for score, letters := range reverseScores() {
		for _, letter := range letters {
			out[letter] = score
		}
	}
	return out
}

// Score a word of scrabble
func Score(word string) int {
	scores := scores()
	score := 0
	for _, char := range strings.ToUpper(word) {
		score += scores[rune(char)]
	}
	return score
}
