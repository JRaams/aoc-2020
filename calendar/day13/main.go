package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)
	earliestDepartmentTime, _ := strconv.Atoi(inputValues[0])
	busIds := loadBusIds(inputValues[1])

	// Part 1
	firstBusId, idleTime := getEarliestBusRide(earliestDepartmentTime, busIds)
	firstBusIdIdleTimeProduct := firstBusId * idleTime
	fmt.Printf("Solution day 13 part 1: bus id multiplied by the number of idle minutes: %d", firstBusIdIdleTimeProduct)
	fmt.Println()
}

func loadBusIds(s string) []int {
	var ids []int
	for _, strId := range strings.Split(s, ",") {
		if strId == "x" {
			continue
		}
		id, _ := strconv.Atoi(strId)
		ids = append(ids, id)
	}
	return ids
}

func getEarliestBusRide(earliestDepartmentTime int, busIds []int) (firstBusId int, idleTime int) {
	shortestIdleTime := 10000000

	for _, busId := range busIds {
		nextDepartmentTime := earliestDepartmentTime + (busId - (earliestDepartmentTime % busId))
		idleTime := nextDepartmentTime - earliestDepartmentTime
		if idleTime < shortestIdleTime {
			shortestIdleTime = idleTime
			firstBusId = busId
		}
	}

	return firstBusId, shortestIdleTime
}
