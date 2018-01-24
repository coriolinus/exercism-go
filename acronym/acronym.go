// Package acronym creates acronyms from strings
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate Create an initialism from the given string
func Abbreviate(s string) string {
	return strings.ToUpper(strings.Join(initial(split(s)), ""))
}

// Split the input string by non-letter characters,
// retaining non-empty values.
func split(s string) []string {
	// first generate the bounds
	var prevBound int
	var slices []string

	for idx, ch := range s {
		if !unicode.IsLetter(ch) {
			if idx != prevBound {
				slices = append(slices, s[prevBound:idx])
			}
			prevBound = idx + 1
		}
	}
	// don't forget to add the terminating slice once we've run out
	// of new characters
	if len(s) != prevBound {
		slices = append(slices, s[prevBound:])
	}

	return slices
}

// Get the first letter of each of a list of strings
func initial(strings []string) []string {
	initials := make([]string, 0, len(strings))
	for _, s := range strings {
		initials = append(initials, s[0:1])
	}
	return initials
}
