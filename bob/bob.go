/*
	Implement "bob" exercise
*/

package bob

import (
	"strings"
	"unicode"
)

const answer = "Sure."
const chill = "Whoa, chill out!"
const stressAnswer = "Calm down, I know what I'm doing!"
const unprompted = "Fine. Be that way!"
const whatever = "Whatever."

// Hey (remark): Respond like a surly teenager
func Hey(remark string) string {
	// I'd prefer to use a `switch` statement here, the way I'd use
	// `match` in Rust, but apparently the syntax doesn't allow that.
	// Therefore, we chain ifs
	remark = strings.TrimSpace(remark)
	if isEmpty(remark) {
		return unprompted
	}
	question := isQuestion(remark)
	yell := isYell(remark)
	if question && yell {
		return stressAnswer
	} else if yell {
		return chill
	} else if question {
		return answer
	}
	return whatever
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isYell(remark string) bool {
	hasLetter := false
	for _, ch := range remark {
		if unicode.IsLetter(ch) {
			hasLetter = true
			if !unicode.IsUpper(ch) {
				return false
			}
		}
	}
	return hasLetter
}

func isEmpty(remark string) bool {
	return len(remark) == 0
}
