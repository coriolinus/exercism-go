package raindrops

import (
	"strconv"
)

// Convert a number to a rain sound
func Convert(n int) string {
	out := ""
	if n%3 == 0 {
		out += "Pling"
	}
	if n%5 == 0 {
		out += "Plang"
	}
	if n%7 == 0 {
		out += "Plong"
	}
	if len(out) == 0 {
		out = strconv.Itoa(n)
	}
	return out
}
