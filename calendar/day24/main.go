package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	tiles := loadTiles(lines)

	// Part a
	a := solveA(tiles)
	fmt.Printf("Solution day 24 part a: %d\n", a)
}

func solveA(t tiles) int {
	helpers.Measure(time.Now(), "")

	blackTiles := 0
	for _, flips := range t {
		if flips%2 == 1 {
			blackTiles++
		}
	}
	return blackTiles
}
