package main

import (
	"fmt"

	"github.com/jraams/aoc-2020/helpers"
)

type waypoint struct {
	x int
	y int
}

func (w *waypoint) followInstruction(i instruction) (shipDx int, shipDy int) {
	switch i.action {
	case "N":
		{
			w.y += i.value
			return 0, 0
		}
	case "S":
		{
			w.y -= i.value
			return 0, 0
		}
	case "E":
		{
			w.x += i.value
			return 0, 0
		}
	case "W":
		{
			w.x -= i.value
			return 0, 0
		}
	case "L":
		{
			w.x, w.y = helpers.RotateCartesianIntCoordsByDegrees(w.x, w.y, -i.value)
			return 0, 0
		}
	case "R":
		{
			w.x, w.y = helpers.RotateCartesianIntCoordsByDegrees(w.x, w.y, i.value)
			return 0, 0
		}
	case "F":
		{
			return i.value * w.x, i.value * w.y
		}
	default:
		{
			panic(fmt.Sprintf("Unknown instruction action: %s", i.action))
		}
	}
}
