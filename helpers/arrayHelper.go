package helpers

import (
	"strconv"
)

// TranslateStringArrToIntArr translates an array of strings to an array of ints
func TranslateStringArrToIntArr(a []string) (c []int) {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i], _ = strconv.Atoi(a[i])
	}
	return b
}

// SumIntArrValues sums up all values in an Int array
func SumIntArrValues(a []int) (b int) {
	b = 0
	for i := 0; i < len(a); i++ {
		b += a[i]
	}
	return b
}

// MultIntArrValues returns the product of all values in Int array
func MultIntArrValues(a []int) (b int) {
	b = 1
	for i := 0; i < len(a); i++ {
		b *= a[i]
	}
	return b
}
