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
	fmt.Printf("Solution part 1: %d bag colors eventually contain at least one shiny gold bag", amountOfBagsWithShinyGoldChild)
	fmt.Println()

	// Part 2
	shinyGoldParentBag := findBagByName(parentBags, "shiny gold")
	totalAmountOfRequiredBags := getTotalBagSum(*shinyGoldParentBag) - 1 // Remove 'shiny gold' parent bag from the count
	fmt.Printf("Solution part 2: %d individual bags are required inside the single shiny gold bag", totalAmountOfRequiredBags)
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
