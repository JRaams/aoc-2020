package main

import (
	"fmt"
	"path/filepath"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)
	intValues := helpers.TranslateStringArrToIntArr(inputValues)

	// Part 1
	firstXMASWeakness := findFirstInvalidNumber(intValues, 25)
	fmt.Printf("Solution part 1: first XMAS weakness number: %d", firstXMASWeakness)
	fmt.Println()
}

func findFirstInvalidNumber(intValues []int, preamble int) int {
	for i := preamble; i < len(intValues)-1; i++ {
		previousNumbers := intValues[i-preamble : i]
		_, err := helpers.GetNEntriesThatMatchX(previousNumbers, 2, intValues[i])
		if err != nil {
			return intValues[i]
		}
	}

	panic("Couldn't find an invalid number")
}
