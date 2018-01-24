package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram if no letter in the string appears more than once
func IsIsogram(s string) bool {
	letters := make(map[rune]uint)
	for _, ch := range strings.ToLower(s) {
		if unicode.IsLetter(ch) {
			letters[ch]++
		}
	}
	for _, count := range letters {
		if count > 1 {
			return false
		}
	}
	return true
}
