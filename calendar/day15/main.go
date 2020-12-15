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
	inputValues := helpers.GetInputValues(inputPath)
	intValues := funk.Map(strings.Split(inputValues[0], ","), func(s string) int {
		intVal, _ := strconv.Atoi(s)
		return intVal
	}).([]int)

	// Part a
	a := solveA(intValues)
	fmt.Printf("Solution day 15 part a: %d\n", a)
}

func solveA(intValues []int) int {
	previousNumbers := map[int][]int{}
	turn, lastNumber := 1, -1

	// Add all starting numbers to 'previousNumbers' map
	for _, n := range intValues {
		previousNumbers[n] = []int{turn}
		lastNumber = n
		turn++
	}

	// Loop until turn 2020
	for turn <= 2020 {
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
