package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	helpers "github.com/jraams/aoc-2020/helpers"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("input file not found...")
	}

	entries := strings.Split(string(input), "\n")
	intEntries := helpers.TranslateStringArrToIntArr(entries)

	// Part 1
	matchingEntries1 := getNEntriesThatMatchX(intEntries, 2, 2020)
	answer1 := helpers.MultIntArrValues(matchingEntries1)
	fmt.Printf("Solution for part 1: %d using values: %v", answer1, matchingEntries1)
	fmt.Println()

	// Part 2
	matchingEntries2 := getNEntriesThatMatchX(intEntries, 3, 2020)
	answer2 := helpers.MultIntArrValues(matchingEntries2)
	fmt.Printf("Solution for part 2: %d using values: %v", answer2, matchingEntries2)
	fmt.Println()
}

func getNEntriesThatMatchX(entries []int, N int, X int) (matchingEntries []int) {
	obj := helpers.CombinationGenerator(entries, N)
	for obj.HasNext() {
		nextEntries := obj.Next()
		if helpers.SumIntArrValues(nextEntries) == X {
			return nextEntries
		}
	}
	panic(fmt.Errorf("No %d entries found that sum up to be %d", N, X))
}
