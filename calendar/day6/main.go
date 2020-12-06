package main

import (
	"fmt"
	"path/filepath"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

type group struct {
	persons []person
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
		countSum += getGroupAnswerCount(group, false)
	}
	fmt.Printf("Solution part 1: countSum is %d", countSum)
	fmt.Println()

	// Part 2
	countSum = 0
	for _, group := range groups {
		countSum += getGroupAnswerCount(group, true)
	}
	fmt.Printf("Solution part 2: countSum is %d", countSum)
	fmt.Println()
}

func getGroups(inputValues []string) []group {
	var groups []group

	currentGroup := group{
		persons: []person{},
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
		}

		currentGroup.persons = append(currentGroup.persons, person)
	}
	groups = append(groups, currentGroup)

	return groups
}

func getGroupAnswerCount(group group, requireAll bool) int {
	answerMap := make(map[string]int)

	for _, person := range group.persons {
		for _, answer := range person.answers {
			if !funk.Contains(answerMap, answer) {
				answerMap[answer] = 0
			}
			answerMap[answer]++
		}
	}

	// If only 1 person needs to give an answer for it to count, just return the length of the answerMap
	if !requireAll {
		return len(answerMap)
	}

	// If ALL people in the group need to have the answer, check if the amount in the answerMap for that answer is equal to the amount of people in the group
	omnipresentAnswerCount := 0
	for _, count := range answerMap {
		if count == len(group.persons) {
			omnipresentAnswerCount++
		}
	}
	return omnipresentAnswerCount
}
