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
	firstInvalidNumber := findFirstInvalidNumber(intValues, 25)
	fmt.Printf("Solution part 1: first XMAS weakness number: %d", firstInvalidNumber)
	fmt.Println()

	// Part 2
	encryptionWeakness := findXMASWeakness(intValues, firstInvalidNumber)
	fmt.Printf("Solution part 2: XMAS encryption weakness: %d", encryptionWeakness)
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

func findXMASWeakness(intValues []int, invalidNumber int) int {
	for i := 0; i < len(intValues)-1; i++ {
		for j := i; j < len(intValues)-1; j++ {
			values := intValues[i:j]
			sum := helpers.SumIntArrValues(values)
			if sum == invalidNumber {
				sort.Ints(values)
				return values[0] + values[len(values)-1]
			}
		}
	}

	panic("Couldn't find XMAS weakness")
}
