package reverse

// String gnirtS
//
// https://stackoverflow.com/a/10030772/504550
func String(s string) string {
	runes := []rune(s)
	for index, complement := 0, len(runes)-1; index < complement; index, complement = index+1, complement-1 {
		runes[index], runes[complement] = runes[complement], runes[index]
	}
	return string(runes)
}
