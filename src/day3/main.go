package main

import (
	"fmt"
	"path/filepath"

	"github.com/jraams/aoc-2020/helpers"
)

type move struct {
	dx int
	dy int
}

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)

	// Part 1
	encounteredTrees, _ := slideDown(inputValues, []move{{dx: 3, dy: 1}})
	fmt.Printf("Solution part 1: encountered %d trees", encounteredTrees)
	fmt.Println()

	// Part 2
	_, productOfEncounteredTrees := slideDown(inputValues, []move{{dx: 1, dy: 1}, {dx: 3, dy: 1}, {dx: 5, dy: 1}, {dx: 7, dy: 1}, {dx: 1, dy: 2}})
	fmt.Printf("Solution part 2: product of encountered trees is %d", productOfEncounteredTrees)
	fmt.Println()
}

func slideDown(input []string, moves []move) (int, int) {
	encounteredTrees, productOfEncounteredTrees := 0, 1
	for _, move := range moves {
		n := slideDownOnce(input, move.dx, move.dy)
		encounteredTrees += n
		productOfEncounteredTrees *= n
	}
	return encounteredTrees, productOfEncounteredTrees
}

func slideDownOnce(input []string, rightVel int, downVel int) int {
	encounteredTrees := 0
	x := 0

	for y := downVel; y < len(input); y += downVel {
		x += rightVel

		square := input[y][x%len(input[y])]
		if string(square) == "#" {
			encounteredTrees++
		}
	}

	return encounteredTrees
}
