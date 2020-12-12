package main

import (
	"fmt"
	"math"
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
			radians := -float64(i.value) * (math.Pi / 180)
			newX := float64(w.x)*math.Cos(radians) + float64(w.y)*math.Sin(radians) // Rotate x by -i.value degrees
			newY := float64(w.y)*math.Cos(radians) - float64(w.x)*math.Sin(radians) // Rotate y by -i.value degrees
			w.x = int(math.Round(newX))
			w.y = int(math.Round(newY))
			return 0, 0
		}
	case "R":
		{
			radians := float64(i.value) * (math.Pi / 180)
			newX := float64(w.x)*math.Cos(radians) + float64(w.y)*math.Sin(radians) // Rotate x by i.value degrees
			newY := float64(w.y)*math.Cos(radians) - float64(w.x)*math.Sin(radians) // Rotate y by i.value degrees
			w.x = int(math.Round(newX))
			w.y = int(math.Round(newY))
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
