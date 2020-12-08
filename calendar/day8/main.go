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
	gameconsole := loadInstructions(inputValues)

	// Part 1
	_, accumulator := runInstructions(gameconsole.clone())
	fmt.Printf("Solution part 1: value in accumulator right before an instruction is executed a second time is: %d", accumulator)
	fmt.Println()

	// Part 2
	isFixed, accumulator := fixInstructions(gameconsole.clone())
	if !isFixed {
		fmt.Println("Couldn't fix instructions for part 2, are you sure you copied the input correctly?")
		return
	}
	fmt.Printf("Solution part 2: value in accumulator after finishing all instructions is: %d", accumulator)
	fmt.Println()
}
