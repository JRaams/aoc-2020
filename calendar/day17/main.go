package main

import (
	"fmt"
	"path/filepath"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)

	// Part a
	a := solveA(lines)
	fmt.Printf("Solution day 17 part a: %d\n", a)

	// Part b
	b := solveB(lines)
	fmt.Printf("Solution day 17 part b: %d\n", b)
}

func solveA(lines []string) int {
	return boot(lines, false)
}

func solveB(lines []string) int {
	return boot(lines, true)
}
