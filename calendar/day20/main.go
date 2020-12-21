package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	tiles := load(lines)

	// Part 1
	fixedTiles := fixTiles(tiles)
	a := solveA(fixedTiles)
	fmt.Printf("Solution day 20 part a: %d\n", a)
}

type tile struct {
	id    int
	image [][]rune
}

// Direction constants
const (
	Top    int = 1
	Right  int = 2
	Bottom int = 3
	Left   int = 4
)

func solveA(fixedTiles [][]*tile) int {
	size := len(fixedTiles) - 1
	topLeft := fixedTiles[0][0]
	topRight := fixedTiles[0][size]
	bottomLeft := fixedTiles[size][0]
	bottomRight := fixedTiles[size][size]

	a := topLeft.id * topRight.id * bottomLeft.id * bottomRight.id
	return a
}

// Bottom: true -> get string value of last row of t.image
// Bottom: false -> get string value of every character in the last column of t.image
func (t *tile) getBorder(dir int) string {
	switch dir {
	// Top
	case 1:
		{
			return string(t.image[0])
		}
		// Right
	case 2:
		{
			border := ""
			for row := 0; row < len(t.image); row++ {
				border += string(t.image[row][len(t.image)-1])
			}
			return border
		}
		// Bottom
	case 3:
		{
			return string(t.image[len(t.image)-1])
		}
		// Left
	case 4:
		{
			border := ""
			for row := 0; row < len(t.image); row++ {
				border += string(t.image[row][0])
			}
			return border
		}
	default:
		{
			panic(fmt.Errorf("getBorder(): unknown direction: %d", dir))
		}
	}
}

func (t *tile) print() {
	for _, row := range t.image {
		fmt.Println(string(row))
	}
}

// Horizontally: true -> reverse all rows
// Horizontally: false -> make a copy of image, for each row, replace it with image[9-index]
func (t *tile) flip(horizontally bool) {
	imageCopy := make([][]rune, len(t.image))
	if horizontally {
		for i := 0; i < len(t.image); i++ {
			imageCopy[i] = helpers.ReverseRuneArr(t.image[i])
		}
	} else {
		for i := 0; i < len(t.image); i++ {
			imageCopy[i] = t.image[len(t.image)-1-i]
		}
	}
	t.image = imageCopy
}

// Rotate the tile 90degs
func (t *tile) rotate() {
	size := len(t.image)
	imageCopy := make([][]rune, size)

	for i := 0; i < size; i++ {
		imageCopy[i] = make([]rune, size)
		for j := 0; j < size; j++ {
			imageCopy[i][j] = t.image[size-j-1][i]
		}
	}
	t.image = imageCopy
}

func (t *tile) bordersMatch(otherT *tile, dir int) bool {
	otherDir := dir + 2
	if otherDir > 4 {
		otherDir -= 4
	}
	return t.getBorder(dir) == otherT.getBorder(otherDir)
}

func load(lines []string) []tile {
	var tiles []tile

	t := tile{}
	for _, line := range lines {
		if len(line) == 0 {
			tiles = append(tiles, t)
			t = tile{}
			continue
		}
		if strings.Contains(line, "Tile") {
			strNum := strings.Replace(strings.Split(line, " ")[1], ":", "", 1)
			t.id = helpers.MustAtoi(strNum)
			t.image = [][]rune{}
			continue
		}
		t.image = append(t.image, []rune(line))
	}
	tiles = append(tiles, t)

	return tiles
}

func fixTiles(allTiles []tile) [][]*tile {
	size := int(math.Floor(math.Sqrt(float64(len(allTiles)))))

	for i := 0; i < len(allTiles); i++ {
		tilesCopy := make([]tile, len(allTiles))
		copy(tilesCopy, allTiles)
		found, tiles := tryTilesetCombinations(size, tilesCopy, tilesCopy[i])
		printTileIds(tiles)
		saveFullImageToFile(tiles)
		if !found {
			continue
		}
		return tiles
	}

	panic("Couldn't find a valid tileset")
}

func printTileIds(a [][]*tile) {
	for _, row := range a {
		s := ""
		for _, tile := range row {
			if tile != nil {
				s += strconv.Itoa(tile.id)
			} else {
				s += "x"
			}
			s += "\t"
		}
		fmt.Println(s)
	}
}

func saveFullImageToFile(tiles [][]*tile) {
	file, err := os.Create("./image")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, tileRow := range tiles {
		for row := 1; row < 9; row++ {
			line := ""
			for _, t := range tileRow {
				for col := 1; col < 9; col++ {
					line += string(t.image[row][col])
				}
			}
			file.WriteString(line + "\n")
		}
	}
}

func tryTilesetCombinations(size int, allTiles []tile, t tile) (bool, [][]*tile) {
	valid, tiles := tryTileset(size, allTiles, t)
	if valid {
		return true, tiles
	}

	for i := 0; i < 4; i++ {
		t.rotate()
		valid, tiles := tryTileset(size, allTiles, t)
		if valid {
			return true, tiles
		}
	}
	t.flip(true)
	for i := 0; i < 4; i++ {
		t.rotate()
		valid, tiles := tryTileset(size, allTiles, t)
		if valid {
			return true, tiles
		}
	}
	t.flip(false)
	for i := 0; i < 4; i++ {
		t.rotate()
		valid, tiles := tryTileset(size, allTiles, t)
		if valid {
			return true, tiles
		}
	}
	t.flip(true)
	for i := 0; i < 4; i++ {
		t.rotate()
		valid, tiles := tryTileset(size, allTiles, t)
		if valid {
			return true, tiles
		}
	}

	return false, nil
}

func tryTileset(size int, allTiles []tile, t tile) (bool, [][]*tile) {
	tiles := make([][]*tile, size)

	for row := 0; row < size; row++ {
		tiles[row] = make([]*tile, size)
	}
	tiles[0][0] = &t

	remainingTiles := make([]tile, len(allTiles))
	copy(remainingTiles, allTiles)
	remainingTiles = funk.Filter(remainingTiles, func(x tile) bool {
		return x.id != t.id
	}).([]tile)

	currentTile := t
	for row := 0; row < size; row++ {
		// Find matching tile for the rest of this row
		for col := 1; col < size; col++ {
			found, t2 := findMatchingTile(remainingTiles, &currentTile, Right)
			if !found {
				return false, tiles
			}

			currentTile = *t2
			tiles[row][col] = t2
			remainingTiles = funk.Filter(remainingTiles, func(x tile) bool {
				return x.id != t2.id
			}).([]tile)
		}

		if len(remainingTiles) == 0 {
			return true, tiles
		}

		// Find matching tile on bottom of currentTile
		found, t2 := findMatchingTile(remainingTiles, tiles[row][0], Bottom)
		if !found {
			return false, tiles
		}
		currentTile = *t2
		tiles[row+1][0] = t2
		remainingTiles = funk.Filter(remainingTiles, func(x tile) bool {
			return x.id != t2.id
		}).([]tile)
	}

	return true, tiles
}

func findMatchingTile(remainingTiles []tile, t *tile, dir int) (bool, *tile) {
	for _, otherTile := range remainingTiles {
		if t.id == otherTile.id {
			continue
		}

		if t.bordersMatch(&otherTile, dir) {
			return true, &otherTile
		}

		for i := 0; i < 4; i++ {
			otherTile.rotate()
			if t.bordersMatch(&otherTile, dir) {
				return true, &otherTile
			}
		}
		otherTile.flip(true)
		for i := 0; i < 4; i++ {
			otherTile.rotate()
			if t.bordersMatch(&otherTile, dir) {
				return true, &otherTile
			}
		}
		otherTile.flip(false)
		for i := 0; i < 4; i++ {
			otherTile.rotate()
			if t.bordersMatch(&otherTile, dir) {
				return true, &otherTile
			}
		}
		otherTile.flip(true)
		for i := 0; i < 4; i++ {
			otherTile.rotate()
			if t.bordersMatch(&otherTile, dir) {
				return true, &otherTile
			}
		}
	}
	return false, nil
}
