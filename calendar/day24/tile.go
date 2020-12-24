package main

import (
	"github.com/thoas/go-funk"
)

// Map x to map of y to 'isBlack' bool
type tileMap map[int]map[int]bool

func (t *tileMap) getBlackTiles() int {
	blackTiles := 0
	for _, x := range *t {
		for _, isBlack := range x {
			if isBlack {
				blackTiles++
			}
		}
	}
	return blackTiles
}

func (t *tileMap) flipTiles() {
	for day := 1; day <= 100; day++ {
		nextState := tileMap{}

		// Find all possible hexagon tile positions
		opts := []pos{}
		for x, yBlackMap := range *t {
			for y := range yBlackMap {
				opts = append(opts, pos{x, y})

				for _, dp := range directions {
					nextX := x + dp.x
					nextY := y + dp.y
					opts = append(opts, pos{x: nextX, y: nextY})
				}
			}
		}

		// > Every day, the tiles are all flipped according to the following rules:
		for _, pos := range opts {
			adj := 0
			for _, dp := range directions {
				if (*t)[pos.x+dp.x] != nil {
					if (*t)[pos.x+dp.x][pos.y+dp.y] {
						adj++
					}
				}
			}

			if nextState[pos.x] == nil {
				nextState[pos.x] = map[int]bool{}
			}

			isBlack := (*t)[pos.x][pos.y]
			// > Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
			if isBlack && (adj == 0 || adj > 2) {
				nextState[pos.x][pos.y] = false

				// > Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
			} else if !isBlack && adj == 2 {
				nextState[pos.x][pos.y] = true
			} else {
				nextState[pos.x][pos.y] = isBlack
			}
		}

		// > The rules are applied simultaneously to every tile; put another way, it is first determined which tiles need
		// > to be flipped, then they are all flipped at the same time.
		*t = nextState
	}
}

func loadTiles(lines []string) tileMap {
	tiles := tileMap{}

	for _, line := range lines {
		pos := getTilePos(line)
		if tiles[pos.x] == nil {
			tiles[pos.x] = map[int]bool{}
		}
		tiles[pos.x][pos.y] = !tiles[pos.x][pos.y]
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
