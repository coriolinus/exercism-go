package diffsquares

func SumOfSquares(until int) int {
	sum := 0
	for n := 0; n <= until; n++ {
		sum += n * n
	}
	return sum
}

func SquareOfSums(until int) int {
	sum := 0
	for n := 0; n <= until; n++ {
		sum += n
	}
	return sum * sum
}

func Difference(until int) int {
	return SquareOfSums(until) - SumOfSquares(until)
}
