package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)
	seatIds := getAllSeatIds(inputValues)

	// Part 1
	sort.Slice(seatIds, func(i int, j int) bool {
		return seatIds[i] > seatIds[j]
	})
	highestSeatId := seatIds[0]
	fmt.Printf("Solution part 1: highest seat id: %d", highestSeatId)
	fmt.Println()

	// Part 2
	lowestSeatId := seatIds[len(seatIds)-1]
	allSeatIds := helpers.GenInt64Array(lowestSeatId, highestSeatId)
	missingSeatIds := helpers.Int64ArrDifference(allSeatIds, seatIds)
	fmt.Printf("Solution part 2: missing seat id: %d", missingSeatIds[0])
	fmt.Println()
}

func getAllSeatIds(inputValues []string) []int64 {
	var seatIds []int64

	for _, line := range inputValues {
		rowStr := strings.ReplaceAll(strings.ReplaceAll(line[:7], "B", "1"), "F", "0")
		colStr := strings.ReplaceAll(strings.ReplaceAll(line[7:], "R", "1"), "L", "0")

		row, _ := strconv.ParseInt(rowStr, 2, 64)
		col, _ := strconv.ParseInt(colStr, 2, 64)
		seatId := row*8 + col

		seatIds = append(seatIds, seatId)
	}

	return seatIds
}
