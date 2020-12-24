package main

import (
	"fmt"
	"path/filepath"
	"strconv"
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

type tiles map[string]int

func loadTiles(lines []string) tiles {
	tiles := tiles{}

	for _, line := range lines {
		x, y := getTilePos(line)
		pos := strconv.Itoa(x) + "," + strconv.Itoa(y)
		tiles[pos]++
	}

	return tiles
}

func getTilePos(line string) (int, int) {
	x, y := 0, 0

	dirs := []byte(line)

	for l := len(dirs); l > 0; {
		c := dirs[0]
		c2 := byte('q')

		if len(dirs) > 1 {
			c2 = dirs[1]
		}

		spliceAmount := 1
		if c == 's' && c2 == 'e' {
			spliceAmount = 2
			y++
		} else if c == 's' && c2 == 'w' {
			spliceAmount = 2
			x--
			y++
		} else if c == 'n' && c2 == 'e' {
			spliceAmount = 2
			x++
			y--
		} else if c == 'n' && c2 == 'w' {
			spliceAmount = 2
			y--
		} else if c == 'e' {
			x++
		} else if c == 'w' {
			x--
		}

		dirs = dirs[spliceAmount:]
		l = len(dirs)
	}
	return x, y
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
