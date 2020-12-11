package main

import "fmt"

type grid [][]position

func (g grid) print() {
	for _, row := range g {
		for _, pos := range row {
			fmt.Print(pos.getChar())
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g grid) applyRules() {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			g[y][x].tempIsOccupied = g[y][x].isOccupied
			if g[y][x].isFloor {
				continue
			}

			if g[y][x].checkRule1(g) {
				g[y][x].tempIsOccupied = true
				continue
			}

			if g[y][x].checkRule2(g) {
				g[y][x].tempIsOccupied = false
				continue
			}
		}
	}

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			g[y][x].isOccupied = g[y][x].tempIsOccupied
		}
	}
}

func (g grid) countOccupiedSeats() int {
	occupiedSeats := 0

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if g[y][x].isOccupied {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}

func loadGrid(inputValues []string) grid {
	width := len(inputValues[0])
	var grid grid

	// Load all positions into the grid
	for y := 0; y < len(inputValues); y++ {
		grid = append(grid, make([]position, width))
		line := inputValues[y]
		for x := 0; x < len(line); x++ {
			char := string(line[x])
			position := makePosition(string(char), y, x)
			grid[y][x] = position
		}
	}

	// Set adjacent seats beforehand
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].isFloor {
				continue
			}
			neighBourSeats := getNeighBourSeats(grid, y, x)
			grid[y][x].adjacentSeats = append(grid[y][x].adjacentSeats, neighBourSeats...)
		}
	}

	return grid
}
