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
	entry1, entry2, err := getTwoEntriesThatMatchX(intEntries, 2020)
	if err != nil {
		panic(fmt.Sprintf("Error getting two entries that match: %s", err.Error()))
	}

	answer1 := entry1 * entry2
	fmt.Printf("Solution found: entry1: %d, entry2: %d, answer: %d", entry1, entry2, answer1)
	fmt.Println()
}

func getTwoEntriesThatMatchX(entries []int, X int) (entry1 int, entry2 int, e error) {
	for i := 0; i < len(entries); i++ {
		entry := entries[i]

		for j := 0; j < len(entries); j++ {
			secondEntry := entries[j]

			if entry+secondEntry == X {
				return entry, secondEntry, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("No entries found that sum up to be %d", X)
}
