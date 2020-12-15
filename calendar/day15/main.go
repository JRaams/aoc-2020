package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	input := loadInput(lines[0])

	// Part a
	a := findLastSpokenNumber(input, 2020)
	fmt.Printf("Solution day 15 part a: %d\n", a)

	// Part b
	b := findLastSpokenNumber(input, 30000000)
	fmt.Printf("Solution day 15 part b: %d\n", b)
}

func loadInput(strInput string) []int {
	return funk.Map(strings.Split(strInput, ","), func(s string) int {
		intVal, _ := strconv.Atoi(s)
		return intVal
	}).([]int)
}

func findLastSpokenNumber(intValues []int, N int) int {
	previousNumbers := map[int][]int{}
	turn, lastNumber := 1, -1

	// Add all starting numbers to 'previousNumbers' map
	for _, n := range intValues {
		previousNumbers[n] = []int{turn}
		lastNumber = n
		turn++
	}

	// Loop until turn N
	for turn <= N {
		speakIndexes := previousNumbers[lastNumber]
		if len(speakIndexes) == 1 {
			lastNumber = 0
		} else {
			lastNumber = speakIndexes[len(speakIndexes)-1] - speakIndexes[len(speakIndexes)-2]
		}

		previousNumbers[lastNumber] = append(previousNumbers[lastNumber], turn)
		turn++
	}

	return lastNumber
}
