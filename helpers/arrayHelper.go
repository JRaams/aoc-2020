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

// GenInt64Array takes a 'min' and 'max' value and creates an array of int64's with the values between them
func GenInt64Array(min int64, max int64) []int64 {
	var a []int64
	for i := min; i <= max; i++ {
		a = append(a, i)
	}
	return a
}

// Int64ArrDifference takes two int64 arrays 'a' and 'b', and returns the difference between them
func Int64ArrDifference(a []int64, b []int64) []int64 {
	m := make(map[int64]bool)
	var diff []int64

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}
