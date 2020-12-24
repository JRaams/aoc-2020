package main

import (
	"github.com/thoas/go-funk"
)

type tiles map[pos]int

func loadTiles(lines []string) tiles {
	tiles := tiles{}

	for _, line := range lines {
		pos := getTilePos(line)
		tiles[pos]++
	}
	return tiles
}

type pos struct {
	x int
	y int
}

var directions = map[string]pos{
	"se": {x: 0, y: 1},
	"sw": {x: -1, y: 1},
	"ne": {x: 1, y: -1},
	"nw": {x: 0, y: -1},
	"e":  {x: 1, y: 0},
	"w":  {x: -1, y: 0},
}

func getTilePos(line string) pos {
	x, y := 0, 0
	dirs := []byte(line)

	for l := len(dirs); l > 0; l = len(dirs) {
		spliceAmount := 2
		if funk.Contains(directions, string(dirs[0])) {
			spliceAmount = 1
		}

		dir := directions[string(dirs[0:spliceAmount])]
		x += dir.x
		y += dir.y

		dirs = dirs[spliceAmount:]
	}
	return pos{x, y}
}
