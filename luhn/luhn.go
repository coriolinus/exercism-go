package luhn

import (
	"strconv"
	"strings"
)

// Valid if the input is a luhn number
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)

	digits := make([]int, 0, len(input))
	for _, ch := range input {
		n, err := strconv.Atoi(string(ch))
		if err != nil {
			return false
		}
		digits = append(digits, n)
	}

	if len(digits) < 2 {
		return false
	}

	rightEven := false
	for index := len(digits) - 1; index >= 0; index-- {
		if rightEven {
			digits[index] = 2 * digits[index]
			if digits[index] > 9 {
				digits[index] -= 9
			}
		}
		rightEven = !rightEven
	}

	sum := 0
	for _, d := range digits {
		sum += d
	}
	return sum%10 == 0
}
