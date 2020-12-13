package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)
	earliestDepartmentTime, _ := strconv.Atoi(inputValues[0])
	busses := loadBusses(inputValues[1])

	// Part 1
	firstBus, idleTime := getEarliestBusRide(earliestDepartmentTime, busses)
	firstBusIdIdleTimeProduct := firstBus.id * idleTime
	fmt.Printf("Solution day 13 part 1: bus id multiplied by the number of idle minutes: %d", firstBusIdIdleTimeProduct)
	fmt.Println()

	// Part 2
	firstCommonTimestamp := getFirstCommonTimestamp(busses)
	fmt.Printf("Solution day 13 part 2: first common timestamp: %d", firstCommonTimestamp)
	fmt.Println()
}
