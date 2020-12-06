package main

import (
	"fmt"
	"path/filepath"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

type group struct {
	persons []person
	answers []string
}

type person struct {
	answers []string
}

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)
	groups := getGroups(inputValues)

	// Part 1
	countSum := 0
	for _, group := range groups {
		countSum += len(group.answers)
	}

	fmt.Printf("Solution part 1: countSum is %d", countSum)
	fmt.Println()
}

func getGroups(inputValues []string) []group {
	var groups []group

	currentGroup := group{
		persons: []person{},
		answers: []string{},
	}
	for _, line := range inputValues {
		if len(line) == 0 {
			groups = append(groups, currentGroup)
			currentGroup = group{
				persons: []person{},
			}
			continue
		}

		person := person{
			answers: []string{},
		}

		for _, c := range line {
			person.answers = append(person.answers, string(c))
			if !funk.Contains(currentGroup.answers, string(c)) {
				currentGroup.answers = append(currentGroup.answers, string(c))
			}
		}

		currentGroup.persons = append(currentGroup.persons, person)
	}
	groups = append(groups, currentGroup)

	return groups
}
