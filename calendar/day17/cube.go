package main

import (
	"errors"
)

var cubes []*cube

type cube struct {
	id         int
	x          int
	y          int
	z          int
	active     bool
	nextState  bool
	neighbours []*cube
}

func loadCubes(lines []string, cycles int) {
	id := 0

	// Create a 3d grid of inactive cubes len(lines) + cycles wide & high, cycles deep since z starts off as 0
	for x := -cycles; x < len(lines)+cycles; x++ {
		for y := -cycles; y < len(lines)+cycles; y++ {
			for z := -cycles; z < cycles+1; z++ {
				cubes = append(cubes, &cube{
					id:         id,
					x:          x,
					y:          y,
					z:          z,
					active:     false,
					neighbours: []*cube{},
				})
				id++
			}
		}
	}

	// Set all neighbours
	for x := -cycles; x < len(lines)+cycles; x++ {
		for y := -cycles; y < len(lines)+cycles; y++ {
			for z := -cycles; z < cycles+1; z++ {
				cube, _ := getCubeAt(x, y, z)
				neighBours := cube.getNeighbours()
				cube.neighbours = neighBours
			}
		}
	}

	// Set initial active cubes
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			r := line[x]
			active := r == '#'
			cube, _ := getCubeAt(x, y, 0)
			cube.active = active
		}
	}
}

func getCubeAt(x int, y int, z int) (*cube, error) {
	for _, cube := range cubes {
		if cube.x == x && cube.y == y && cube.z == z {
			return cube, nil
		}
	}
	return nil, errors.New("No cube found")
}

func (c *cube) getNeighbours() []*cube {
	var neighBours []*cube
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				otherCube, err := getCubeAt(c.x+dx, c.y+dy, c.z+dz)
				if err != nil || c.id == otherCube.id {
					continue
				}
				neighBours = append(neighBours, otherCube)
			}
		}
	}
	return neighBours
}

func (c *cube) getActiveNeighbourCount() int {
	active := 0
	for _, neighBour := range c.neighbours {
		if neighBour.active {
			active++
		}
	}
	return active
}

func (c *cube) applyRules() {
	activeCount := c.getActiveNeighbourCount()
	c.nextState = false

	if c.active {
		if activeCount == 2 || activeCount == 3 {
			c.nextState = true
		}
	} else {
		if activeCount == 3 {
			c.nextState = true
		}
	}
}
