package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)
	intValues := helpers.TranslateStringArrToIntArr(inputValues)

	// Part 1
	joltDifferenceMap := getJoltDifferences(intValues)
	productOfD1D3 := joltDifferenceMap[1] * joltDifferenceMap[3]
	fmt.Printf("Solution part 1: product of 1-jolt and 3-jolt differences is: %d", productOfD1D3)
	fmt.Println()

	// Part 2
	arrangements := getProdAdapterArrangements(intValues)
	fmt.Printf("Solution part 2: product of all distinct arrangements of adapters: %d", arrangements)
	fmt.Println()
}

func getJoltDifferences(intValues []int) map[int]int {
	sort.Ints(intValues)
	joltDifferences := map[int]int{1: 1, 3: 1}

	for i := 1; i < len(intValues); i++ {
		d := intValues[i] - intValues[i-1]
		joltDifferences[d]++
	}
	return joltDifferences
}

func getProdAdapterArrangements(intValues []int) int {
	intValues = append([]int{0}, intValues...)
	intValues = append(intValues, intValues[len(intValues)-1]+3)

	combinationMap := []int{1, 2, 4, 7}
	arrangementProduct := 1

	adapterCount := 0
	for i := 1; i+1 < len(intValues); i++ {
		dLeft := intValues[i] - intValues[i-1]
		dRight := intValues[i+1] - intValues[i]
		if dLeft == 1 && dRight == 1 {
			adapterCount++
			continue
		}
		arrangementProduct *= combinationMap[adapterCount]
		adapterCount = 0
	}
	return arrangementProduct
}
