package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

type passport map[string]string

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

func isPassportValid(passport passport, betterValidation bool) bool {
	for _, requiredField := range requiredFields {
		if !funk.Contains(passport, requiredField) {
			return false
		}
	}

	if !betterValidation {
		return true
	}

	checks := []bool{
		validateBirthYear(passport["byr"]),
		validateIssueYear(passport["iyr"]),
		validateExpirationYear(passport["eyr"]),
		validateHeight(passport["hgt"]),
		validateHairColor(passport["hcl"]),
		validateEyeColor(passport["ecl"]),
		validatePassportId(passport["pid"]),
	}
	return helpers.All(checks, true)
}

func validateBirthYear(byr string) bool {
	birthYear, err := strconv.ParseInt(byr, 10, 64)
	return err == nil && len(byr) == 4 && 1920 <= birthYear && birthYear <= 2002
}

func validateIssueYear(iyr string) bool {
	issueYear, err := strconv.ParseInt(iyr, 10, 64)
	return err == nil && len(iyr) == 4 && 2010 <= issueYear && issueYear <= 2020
}

func validateExpirationYear(eyr string) bool {
	expYear, err := strconv.ParseInt(eyr, 10, 64)
	return err == nil && len(eyr) == 4 && 2020 <= expYear && expYear <= 2030
}

func validateHeight(hgt string) bool {
	heightAmount, err := strconv.ParseInt(hgt[:len(hgt)-2], 10, 64)
	if err != nil {
		return false
	}
	heightMeasurement := hgt[len(hgt)-2:]
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
	return true
}

func validateHairColor(hcl string) bool {
	matched, err := regexp.MatchString("^#[0-9a-z]{6}$", hcl)
	return err == nil && matched
}

var allowedEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func validateEyeColor(ecl string) bool {
	return funk.Contains(allowedEyeColors, ecl)
}

func validatePassportId(pid string) bool {
	matched, err := regexp.MatchString("^[0-9]{9}$", pid)
	return err == nil && matched
}
