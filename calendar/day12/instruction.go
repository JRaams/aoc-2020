package main

import (
	"fmt"
	"math"
	"strconv"
)

type instruction struct {
	action string
	value  int
}

func (i instruction) followInstruction(s ship) (x int, y int, dir int) {
	switch i.action {
	case "N":
		{
			return 0, i.value, 0
		}
	case "S":
		{
			return 0, -i.value, 0
		}
	case "E":
		{
			return i.value, 0, 0
		}
	case "W":
		{
			return -i.value, 0, 0
		}
	case "L":
		{
			return 0, 0, -i.value
		}
	case "R":
		{
			return 0, 0, i.value
		}
	case "F":
		{
			radians := float64(s.dir) * (math.Pi / 180)
			return int(math.Sin(radians)) * i.value, int(math.Cos(radians)) * i.value, 0
		}
	default:
		{
			panic(fmt.Sprintf("Unknown action: %s", i.action))
		}
	}
}

func loadInstructions(inputValues []string) []instruction {
	var instructions []instruction

	for _, line := range inputValues {
		action := string(line[0])
		value, _ := strconv.Atoi(line[1:])

		instructions = append(instructions, instruction{
			action: action,
			value:  value,
		})
	}

	return instructions
}
