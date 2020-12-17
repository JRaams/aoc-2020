package main

import (
	"errors"
	"fmt"
)

const cycleCount = 6

var cubes map[string]*cube

type pos struct {
	x int
	y int
	w int
	z int
}

type cube struct {
	id         int
	pos        pos
	active     bool
	nextState  bool
	neighbours []*cube
}

func loadCubes(lines []string, cycles int, wDim bool) {
	id := 0
	cubes = make(map[string]*cube)

	// Create a 3d/4d grid of inactive cubes len(lines) + cycles wide & high, cycles deep since z starts off as 0
	for x := -cycles; x < len(lines)+cycles; x++ {
		for y := -cycles; y < len(lines)+cycles; y++ {
			for z := -cycles; z < cycles+1; z++ {
				if wDim {
					for w := -cycles; w < cycles+1; w++ {
						c := &cube{
							id: id,
							pos: pos{
								x: x,
								y: y,
								z: z,
								w: w,
							},
							active:     false,
							neighbours: []*cube{},
						}
						k := getKey(x, y, z, w)
						cubes[k] = c
						id++
					}
				} else {
					c := &cube{
						id: id,
						pos: pos{
							x: x,
							y: y,
							z: z,
							w: 0,
						},
						active:     false,
						neighbours: []*cube{},
					}
					k := getKey(x, y, z, 0)
					cubes[k] = c
					id++
				}
			}
		}
	}

	// Set all neighbours
	for x := -cycles; x < len(lines)+cycles; x++ {
		for y := -cycles; y < len(lines)+cycles; y++ {
			for z := -cycles; z < cycles+1; z++ {
				if wDim {
					for w := -cycles; w < cycles+1; w++ {
						cube, _ := getCubeAt(x, y, z, w)
						neighBours := cube.getNeighbours(true)
						cube.neighbours = neighBours
					}
				} else {
					cube, _ := getCubeAt(x, y, z, 0)
					neighBours := cube.getNeighbours(false)
					cube.neighbours = neighBours
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

func boot(lines []string, wDim bool) int {
	loadCubes(lines, cycleCount, wDim)
	for cycle := 0; cycle < cycleCount; cycle++ {
		for _, cube := range cubes {
			cube.applyRules()
		}
		for _, cube := range cubes {
			cube.active = cube.nextState
		}
	}

	activeCubes := 0
	for _, c := range cubes {
		if c.active {
			activeCubes++
		}
	}
	return activeCubes
}

func getKey(x int, y int, z int, w int) string {
	return fmt.Sprintf("%d-%d-%d-%d", x, y, z, w)
}

func getCubeAt(x int, y int, z int, w int) (*cube, error) {
	k := getKey(x, y, z, w)
	if cube, ok := cubes[k]; ok {
		return cube, nil
	}
	return nil, errors.New("No cube found")
}

func (c *cube) getNeighbours(wDim bool) []*cube {
	var neighBours []*cube
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if wDim {
					for dw := -1; dw <= 1; dw++ {
						otherCube, err := getCubeAt(c.pos.x+dx, c.pos.y+dy, c.pos.z+dz, c.pos.w+dw)
						if err != nil || c.id == otherCube.id {
							continue
						}
						neighBours = append(neighBours, otherCube)
					}
				} else {
					otherCube, err := getCubeAt(c.pos.x+dx, c.pos.y+dy, c.pos.z+dz, 0)
					if err != nil || c.id == otherCube.id {
						continue
					}
					neighBours = append(neighBours, otherCube)
				}
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
