package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"

	"github.com/jraams/aoc-2020/helpers"
)

type passport map[string]string

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

func loadPassports(inputValues []string) []passport {
	var passports []passport

	currentPassport := passport{}
	for _, line := range inputValues {
		// Empty line -> move on to next passport
		if len(line) == 0 {
			passports = append(passports, currentPassport)
			currentPassport = passport{}
			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			keyValue := strings.Split(field, ":")
			currentPassport[keyValue[0]] = keyValue[1]
		}
	}
	passports = append(passports, currentPassport)

	return passports
}

func validatePassports(passports []passport, betterValidation bool) int {
	valid := 0

	for _, passport := range passports {
		if isPassportValid(passport, betterValidation) {
			valid++
		}
	}

	return valid
}

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var allowedEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

//gocyclo:ignore
func isPassportValid(passport passport, betterValidation bool) bool {
	// Check that required fields are present
	for _, requiredField := range requiredFields {
		if !funk.Contains(passport, requiredField) {
			return false
		}
	}

	if !betterValidation {
		return true
	}

	// Check birth year
	birthYear, err := strconv.ParseInt(passport["byr"], 10, 64)
	if err != nil || len(passport["byr"]) != 4 || birthYear < 1920 || birthYear > 2002 {
		return false
	}

	// Check issue year
	issueYear, err := strconv.ParseInt(passport["iyr"], 10, 64)
	if err != nil || len(passport["iyr"]) != 4 || issueYear < 2010 || issueYear > 2020 {
		return false
	}

	// Check expiration year
	expYear, err := strconv.ParseInt(passport["eyr"], 10, 64)
	if err != nil || len(passport["eyr"]) != 4 || expYear < 2020 || expYear > 2030 {
		return false
	}

	// Check height
	heightAmount, err := strconv.ParseInt(passport["hgt"][:len(passport["hgt"])-2], 10, 64)
	if err != nil {
		return false
	}
	heightMeasurement := passport["hgt"][len(passport["hgt"])-2:]
	if heightMeasurement == "cm" {
		if heightAmount < 150 || heightAmount > 193 {
			return false
		}
	} else if heightMeasurement == "in" {
		if heightAmount < 59 || heightAmount > 76 {
			return false
		}
	} else {
		return false
	}

	// Check hair color
	matched, err := regexp.MatchString("^#[0-9a-z]{6}$", passport["hcl"])
	if err != nil || !matched {
		return false
	}

	// Check eye color
	if !funk.Contains(allowedEyeColors, passport["ecl"]) {
		return false
	}

	// Check passport id
	matched, err = regexp.MatchString("^[0-9]{9}$", passport["pid"])
	if err != nil || !matched {
		return false
	}

	return true
}
