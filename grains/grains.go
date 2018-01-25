package grains

// SquaresOnBoard of a chess board
const SquaresOnBoard = 64

// Error type
type Error string

func (e Error) Error() string {
	return string(e)
}

// Square returns the number of grains on a given square
func Square(n int) (uint64, error) {
	if n <= 0 {
		return 0, Error("no '0' space on a chess board")
	} else if n > SquaresOnBoard {
		return 0, Error("No squares over 64 on a chess board")
	}
	return 1 << uint(n-1), nil
}

// Total amount of grain on the board
// Precomputed because this function takes no inputs
// and has no side effects.
func Total() uint64 {
	return 0xffffffffffffffff
}
