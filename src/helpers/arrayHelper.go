package helpers

import "strconv"

func TranslateStringArrToIntArr(a []string) (c []int) {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i], _ = strconv.Atoi(a[i])
	}
	return b
}
