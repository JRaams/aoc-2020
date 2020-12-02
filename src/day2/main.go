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
	validEntries := getValidPwentries(pwentries)
	fmt.Printf("Solution part 1: %d passwords are valid!", len(validEntries))
	fmt.Println()
}

type pwentry struct {
	min      int64
	max      int64
	letter   string
	password string
}

func getPwentries(inputValues []string) (pwentries []pwentry) {
	pwentries = make([]pwentry, len(inputValues))

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

func getValidPwentries(pwentries []pwentry) (_validEntries []pwentry) {
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
