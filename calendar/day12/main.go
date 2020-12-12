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
	s1 := ship{dir: 90} // face east
	s1.followInstructions(instructions)
	manhattanDist1 := s1.getManhattanDist()
	fmt.Printf("Solution part 1: the ships manhattan distance from the starting position is: %d", manhattanDist1)
	fmt.Println()

	// Part 2
	waypoint := waypoint{x: 10, y: 1}
	s2 := ship{}
	s2.followWaypoint(&waypoint, instructions)
	manhattanDist2 := s2.getManhattanDist()
	fmt.Printf("Solution part 2: the ships manhattan distance from the starting position is: %d", manhattanDist2)
	fmt.Println()
}
