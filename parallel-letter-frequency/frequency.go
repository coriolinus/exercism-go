package letter

// FreqMap is a map of letter frequencies
type FreqMap map[rune]int

// Frequency counts occurrences of a char in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts occurrences of a char in a string concurrently
func ConcurrentFrequency(ss []string) FreqMap {
	subresults := make(chan FreqMap)

	for _, s := range ss {
		// launch worker goroutine to count chunk frequencies
		go func(s string) {
			subresults <- Frequency(s)
		}(s)
	}

	m := FreqMap{}
	for srIndex := 0; srIndex < len(ss); srIndex++ {
		subresult := <-subresults

		for k, v := range subresult {
			m[k] += v
		}
	}

	return m
}
