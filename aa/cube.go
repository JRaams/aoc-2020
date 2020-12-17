package main

import (
	"errors"
)

var cubes []*cube

type cube struct {
	id        int
	x         int
	y         int
	z         int
	w         int
	active    bool
	nextState bool
}

func loadCubes(lines []string, cycles int, wDim bool) {
	id := 0
	cubes = []*cube{}

	// Create a 3d grid of inactive cubes len(lines) + cycles wide & high, cycles deep since z starts off as 0
	for x := -cycles; x < len(lines)+cycles; x++ {
		for y := -cycles; y < len(lines)+cycles; y++ {
			for z := -cycles; z < cycles+1; z++ {
				if wDim {
					for w := -cycles; w < cycles+1; w++ {
						cubes = append(cubes, &cube{
							id:     id,
							x:      x,
							y:      y,
							z:      z,
							w:      w,
							active: false,
						})
						id++
					}
				} else {
					cubes = append(cubes, &cube{
						id:     id,
						x:      x,
						y:      y,
						z:      z,
						w:      0,
						active: false,
					})
					id++
				}
			}
		}
	}

	// Set initial active cubes
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			r := line[x]
			active := r == '#'
			cube, _ := getCubeAt(x, y, 0, 0)
			cube.active = active
		}
	}
}

func getCubeAt(x int, y int, z int, w int) (*cube, error) {
	for _, cube := range cubes {
		if cube.x == x && cube.y == y && cube.z == z && cube.w == w {
			return cube, nil
		}
	}
	return nil, errors.New("No cube found")
}

func (c *cube) getNeighbours(wDim bool) []cube {
	var neighBours []cube
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if wDim {
					for dw := -1; dw <= 1; dw++ {
						otherCube, err := getCubeAt(c.x+dx, c.y+dy, c.z+dz, c.w+dw)
						if err != nil || c.id == otherCube.id {
							continue
						}
						neighBours = append(neighBours, *otherCube)
					}
				} else {
					otherCube, err := getCubeAt(c.x+dx, c.y+dy, c.z+dz, 0)
					if err != nil || c.id == otherCube.id {
						continue
					}
					neighBours = append(neighBours, *otherCube)
				}
			}
		}
	}
	return neighBours
}

func (c *cube) getNeighbourCount(wDim bool) int {
	active := 0
	for _, neighBour := range c.getNeighbours(wDim) {
		if neighBour.active {
			active++
		}
	}
	return active
}

func (c *cube) applyRules(wDim bool) {
	activeCount := c.getNeighbourCount(wDim)
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
