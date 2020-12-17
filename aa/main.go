package main

import (
	"fmt"
	"path/filepath"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

const CYCLE_COUNT = 6

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)

	// Part a
	loadCubes(lines, CYCLE_COUNT, false)
	a := solveA()
	fmt.Printf("Solution day 17 part a: %d\n", a)

	// // Part b
	// loadCubes(lines, CYCLE_COUNT, true)
	// b := solveB()
	// fmt.Printf("Solution day 17 part b: %d\n", b)
}

func solveA() int {
	for cycle := 0; cycle < CYCLE_COUNT; cycle++ {
		for _, cube := range cubes {
			cube.applyRules(false)
		}
		for _, cube := range cubes {
			cube.active = cube.nextState
		}
	}
	activeCubes := funk.Filter(cubes, func(c *cube) bool {
		return c.active
	}).([]*cube)
	return len(activeCubes)
}

func solveB() int {
	for cycle := 0; cycle < CYCLE_COUNT; cycle++ {
		for _, cube := range cubes {
			cube.applyRules(true)
		}
		for _, cube := range cubes {
			cube.active = cube.nextState
		}
	}
	activeCubes := funk.Filter(cubes, func(c *cube) bool {
		return c.active
	}).([]*cube)
	return len(activeCubes)
}
