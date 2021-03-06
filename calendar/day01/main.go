package main

import (
	"fmt"
	"path/filepath"

	helpers "github.com/jraams/aoc-2020/helpers"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	stringInputValues := helpers.GetInputValues(inputPath)
	intEntries := helpers.TranslateStringArrToIntArr(stringInputValues)

	// Part 1
	matchingEntries1, _ := helpers.GetNEntriesThatMatchX(intEntries, 2, 2020)
	answer1 := helpers.MultIntArrValues(matchingEntries1)
	fmt.Printf("Solution for part 1: %d using values: %v", answer1, matchingEntries1)
	fmt.Println()

	// Part 2
	matchingEntries2, _ := helpers.GetNEntriesThatMatchX(intEntries, 3, 2020)
	answer2 := helpers.MultIntArrValues(matchingEntries2)
	fmt.Printf("Solution for part 2: %d using values: %v", answer2, matchingEntries2)
	fmt.Println()
}
