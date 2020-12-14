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

	// Part 1
	program1 := loadProgram(inputValues)
	program1.run(1)
	sumOfAllMemoryValues1 := program1.getSumOfAllValues()
	fmt.Printf("Solution day 14 part 1: sum of all memory values: %d", sumOfAllMemoryValues1)
	fmt.Println()

	// Part 2
	program2 := loadProgram(inputValues)
	program2.run(2)
	sumOfAllMemoryValues2 := program2.getSumOfAllValues()
	fmt.Printf("Solution day 14 part 2: sum of all memory values: %d", sumOfAllMemoryValues2)
	fmt.Println()
}
