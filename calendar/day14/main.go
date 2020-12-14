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
	program := loadProgram(inputValues)

	// Part 1
	program.run()
	sumOfAllMemoryValues := program.getSumOfAllValues()
	fmt.Printf("Solution day 14 part 1: sum of all memory values: %d", sumOfAllMemoryValues)
	fmt.Println()
}
