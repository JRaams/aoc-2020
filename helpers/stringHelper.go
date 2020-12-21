package helpers

// ReverseString takes a string 's' and reverses it
func ReverseString(s string) string {
	runes := []rune(s)
	return string(ReverseRuneArr(runes))
}

// ReverseRuneArr takes a rune array 'runes' and reverses it
func ReverseRuneArr(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}
