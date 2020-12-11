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
	grid := loadGrid(inputValues, true)
	occupiedSeats := grid.simulate(4)
	fmt.Printf("Solution part 1: %d seats are occupied after the chaos stabilizes", occupiedSeats)
	fmt.Println()

	// Part 2
	grid = loadGrid(inputValues, false)
	occupiedSeats = grid.simulate(5)
	fmt.Printf("Solution part 2: %d seats are occupied once equilibrium is reached", occupiedSeats)
	fmt.Println()
}
