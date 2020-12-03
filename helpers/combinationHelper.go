package helpers

// Yoinked from:
// https://leetcode.com/problems/iterator-for-combination/discuss/502469/

type combinationIterator struct {
	Permutations *[][]int
	offset       int
}

// CombinationGenerator generates all possible combinations of length 'combinationLength' from int array 'values'
func CombinationGenerator(values []int, combinationLength int) combinationIterator {
	permutations := make([][]int, 0)
	helper(values, combinationLength, &permutations, make([]int, 0), 0)
	return combinationIterator{Permutations: &permutations, offset: 0}
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

// Get the next iteration of combinations
func (iterator *combinationIterator) Next() []int {
	next := (*iterator.Permutations)[iterator.offset]
	iterator.offset++
	return next
}

// Check if a next interation is available
func (iterator *combinationIterator) HasNext() bool {
	return iterator.offset < len(*iterator.Permutations)
}
