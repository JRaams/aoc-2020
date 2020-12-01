package helpers

import (
	"strconv"
)

func TranslateStringArrToIntArr(a []string) (c []int) {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i], _ = strconv.Atoi(a[i])
	}
	return b
}

func SumIntArrValues(a []int) (b int) {
	b = 0
	for i := 0; i < len(a); i++ {
		b += a[i]
	}
	return b
}

func MultIntArrValues(a []int) (b int) {
	b = 1
	for i := 0; i < len(a); i++ {
		b *= a[i]
	}
	return b
}
