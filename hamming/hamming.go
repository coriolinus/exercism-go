package hamming

// Error type for hamming package
type Error string

func (e Error) Error() string {
	return string(e)
}

// Distance is number of variant bytes between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, Error("Hamming distance only defined for strings of equal length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
