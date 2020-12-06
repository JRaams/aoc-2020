package main

import (
	"fmt"
	"path/filepath"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)
	allPassports := loadPassports(inputValues)

	// Part 1
	amountOfValidPassports := validatePassports(allPassports, false)
	fmt.Printf("Solution part 1: %d valid passports", amountOfValidPassports)
	fmt.Println()

	// Part 2
	amountOfValidPassports = validatePassports(allPassports, true)
	fmt.Printf("Solution part 2: %d valid passports", amountOfValidPassports)
	fmt.Println()
}
