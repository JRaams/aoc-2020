package helpers

// Yoinked from:
// https://leetcode.com/problems/iterator-for-combination/discuss/502469/Go-Beat-96.15-Solution-(with-Explaination)

type CombinationIterator struct {
	Permutations *[][]int
	offset       int
}

func CombinationGenerator(values []int, combinationLength int) CombinationIterator {
	permutations := make([][]int, 0)
	helper(values, combinationLength, &permutations, make([]int, 0), 0)
	return CombinationIterator{Permutations: &permutations, offset: 0}
}

func helper(values []int, l int, per *[][]int, comb []int, start int) {
	if len(comb) == l {
		*per = append(*per, comb)
		return
	}
	for i := start; i < len(values); i++ {
		helper(values, l, per, append(comb, values[i]), i+1)
	}
}

func (this *CombinationIterator) Next() []int {
	next := (*this.Permutations)[this.offset]
	this.offset++
	return next
}

func (this *CombinationIterator) HasNext() bool {
	return this.offset < len(*this.Permutations)
}
