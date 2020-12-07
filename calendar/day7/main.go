package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)
	cleanedInputValues := cleanInputValues(inputValues)
	parentBags := getParentBags(cleanedInputValues)

	// Part 1
	amountOfBagsWithShinyGoldChild := getAmountBagsWithName(parentBags, "shiny gold")
	amountOfBagsWithShinyGoldChild-- // Remove 'shiny gold' parent bag from the count
	fmt.Printf("Solution part 1: %d bag colors eventually contain at least one shiny gold bad", amountOfBagsWithShinyGoldChild)
	fmt.Println()
}

func cleanInputValues(inputValues []string) []string {
	lines := []string{}
	for _, line := range inputValues {
		temp := strings.ReplaceAll(line, " bags", "")
		temp = strings.ReplaceAll(temp, " bag", "")
		temp = temp[:len(temp)-1]
		lines = append(lines, temp)
	}
	return lines
}
