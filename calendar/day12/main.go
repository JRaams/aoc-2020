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
	instructions := loadInstructions(inputValues)

	// Part 1
	ship := ship{dir: 90} // face east
	ship.followInstructions(instructions)
	manhattanDist := ship.getManhattanDist()
	fmt.Printf("Solution part 1: the ships manhattan distance from the starting position is: %d", manhattanDist)
	fmt.Println()
}
