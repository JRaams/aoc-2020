package main

import (
	"fmt"
	"time"

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

		startH := time.Now()
		// Find all possible hexagon tiles
		opts := tileMap{}
		for x, yBlackMap := range *t {
			if opts[x] == nil {
				opts[x] = map[int]bool{}
			}
			for y, isBlack := range yBlackMap {
				opts[x][y] = isBlack

				for _, dp := range directions {
					nextX := x + dp.x
					nextY := y + dp.y
					if opts[nextX] != nil {
						if opts[nextX][nextY] {
							continue
						}
					}
					if opts[nextX] == nil {
						opts[nextX] = map[int]bool{}
					}
					opts[nextX][nextY] = false
				}
			}
		}
		fmt.Printf("Find hex tile: %s\n", time.Since(startH))

		startF := time.Now()
		// > Every day, the tiles are all flipped according to the following rules:
		for x, yBlackMap := range opts {
			for y, isBlack := range yBlackMap {
				adj := 0
				for _, dp := range directions {
					if funk.Contains(*t, x+dp.x) {
						if funk.Contains((*t)[x+dp.x], y+dp.y) {
							if (*t)[x+dp.x][y+dp.y] {
								adj++
							}
						}
					}
				}

				if nextState[x] == nil {
					nextState[x] = map[int]bool{}
				}

				// > Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
				if isBlack && (adj == 0 || adj > 2) {
					nextState[x][y] = false

					// > Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
				} else if !isBlack && adj == 2 {
					nextState[x][y] = true
				} else {
					nextState[x][y] = isBlack
				}
			}
		}
		fmt.Printf("Flip tiles: %s\n", time.Since(startF))

		// > The rules are applied simultaneously to every tile; put another way, it is first determined which tiles need
		// > to be flipped, then they are all flipped at the same time.
		*t = nextState
		fmt.Printf("Day %d: %d\n", day, t.getBlackTiles())
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
