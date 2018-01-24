package twofer

import (
	"fmt"
)

// ShareWith says "One for you, one for me."
func ShareWith(you string) string {
	if len(you) == 0 {
		you = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", you)
}
