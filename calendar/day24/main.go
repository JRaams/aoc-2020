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

	// Part b
	b := solveB(tiles)
	fmt.Printf("Solution day 24 part b: %d\n", b)
}

func solveA(t tileMap) int {
	defer helpers.Measure(time.Now(), "")
	a := t.getBlackTiles()
	return a
}

func solveB(t tileMap) int {
	defer helpers.Measure(time.Now(), "")
	t.flipTiles()
	b := t.getBlackTiles()
	return b
}
