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
	grid := loadGrid(inputValues)

	// Part 1
	occupiedSeats := solveA(grid)
	fmt.Printf("Solution part 1: %d seats are occupied after the chaos stabilizes", occupiedSeats)
	fmt.Println()
}

func solveA(g grid) int {
	oldOccupiedSeats := -1
	for {
		g.applyRules()
		newOccupiedSeats := g.countOccupiedSeats()
		if oldOccupiedSeats == newOccupiedSeats {
			return newOccupiedSeats
		}
		oldOccupiedSeats = newOccupiedSeats
	}
}
