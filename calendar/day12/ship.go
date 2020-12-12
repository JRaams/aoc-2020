package main

import (
	"math"
)

type ship struct {
	x   int // Amount of units horizonally, positive = east, negative = west
	y   int // Amount of units vertically, positive = north, negative = south
	dir int // Direction of the ship in degrees
}

func (s *ship) followInstructions(instructions []instruction) {
	for _, instruction := range instructions {
		x, y, dir := instruction.followInstruction(*s)
		s.dir += dir
		s.x += x
		s.y += y
	}
}

func (s *ship) getManhattanDist() int {
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func (s *ship) followWaypoint(waypoint *waypoint, instructions []instruction) {
	for _, instruction := range instructions {
		shipDx, shipDy := waypoint.followInstruction(instruction)
		s.x += shipDx
		s.y += shipDy
	}
}
