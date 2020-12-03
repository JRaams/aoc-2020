package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)

	pwentries := getPwentries(inputValues)
	validEntriesPart1 := getValidPwentriesPart1(pwentries)
	fmt.Printf("Solution part 1: %d passwords are valid!", len(validEntriesPart1))
	fmt.Println()

	validEntriesPart2 := getValidPwentriesPart2(pwentries)
	fmt.Printf("Solution part 2: %d passwords are valid!", len(validEntriesPart2))
	fmt.Println()
}

type pwentry struct {
	min      int64
	max      int64
	letter   string
	password string
}

func getPwentries(inputValues []string) (_pwentries []pwentry) {
	var pwentries []pwentry

	for i := 0; i < len(inputValues); i++ {
		inputParts := strings.Split(inputValues[i], " ")
		minmax := strings.Split(inputParts[0], "-")
		min, _ := strconv.ParseInt(minmax[0], 10, 64)
		max, _ := strconv.ParseInt(minmax[1], 10, 64)
		letter := inputParts[1][:1]
		password := inputParts[2]

		pwentry := pwentry{min, max, letter, password}
		pwentries = append(pwentries, pwentry)
	}

	return pwentries
}

func getValidPwentriesPart1(pwentries []pwentry) (_validEntries []pwentry) {
	var validEntries []pwentry

	for i := 0; i < len(pwentries); i++ {
		entry := pwentries[i]
		c := int64(strings.Count(entry.password, entry.letter))

		if c >= entry.min && c <= entry.max {
			validEntries = append(validEntries, entry)
		}
	}

	return validEntries
}

func getValidPwentriesPart2(pwentries []pwentry) (_validEntries []pwentry) {
	var validEntries []pwentry

	for i := 0; i < len(pwentries); i++ {
		entry := pwentries[i]

		isFirstPosValid := string(entry.password[entry.min-1]) == entry.letter
		isSecondPosValid := string(entry.password[entry.max-1]) == entry.letter
		if isFirstPosValid && !isSecondPosValid || !isFirstPosValid && isSecondPosValid {
			validEntries = append(validEntries, entry)
		}
	}

	return validEntries
}
