package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	rules, ownTicket, nearbyTickets := loadTicketData(lines)

	// Part a
	a, validNearbyTickets := solveA(nearbyTickets, rules)
	fmt.Printf("Solution day 16 part a: %d\n", a)

	// Part b
	b := solveB(rules, ownTicket, validNearbyTickets)
	fmt.Printf("Solution day 16 part b: %d\n", b)
}

func solveA(nearbyTickets []ticket, rules []rule) (int, []ticket) {
	errorRateSum := 0
	var validNearbyTickets []ticket
	for _, ticket := range nearbyTickets {
		ticketIsValid, errorRate := ticket.isValid(rules)
		if ticketIsValid {
			validNearbyTickets = append(validNearbyTickets, ticket)
		} else {
			errorRateSum += errorRate
		}
	}
	return errorRateSum, validNearbyTickets
}

func solveB(rules []rule, ownTicket ticket, validNearbyTickets []ticket) int {
	// Get all possible rule indexes for each rule
	allPossibleRuleIndexesMap := map[string][]int{}
	for _, rule := range rules {
		possibleRuleIndexes := getPossibleRuleIndexes(rule, append(validNearbyTickets, ownTicket))
		allPossibleRuleIndexesMap[rule.name] = possibleRuleIndexes
	}

	// Find which rules map to what index
	actualRuleNameIndexMap := map[string]int{}
	for len(allPossibleRuleIndexesMap) > 0 {
		for ruleName, possibleIndexes := range allPossibleRuleIndexesMap {
			if len(possibleIndexes) == 1 {
				actualRuleNameIndexMap[ruleName] = possibleIndexes[0]

				for ruleName2 := range allPossibleRuleIndexesMap {
					allPossibleRuleIndexesMap[ruleName2] = funk.FilterInt(allPossibleRuleIndexesMap[ruleName2], func(x int) bool {
						return x != possibleIndexes[0]
					})
				}

				delete(allPossibleRuleIndexesMap, ruleName)
				break
			}
		}
	}

	// Get product of all 6 rules that contain "departure"
	b := 1
	for ruleName, ruleIndex := range actualRuleNameIndexMap {
		if strings.Contains(ruleName, "departure") {
			b *= ownTicket.values[ruleIndex]
		}
	}
	return b
}

func getPossibleRuleIndexes(r rule, tickets []ticket) []int {
	var validIndexes []int
	for i := 0; i < len(tickets[0].values); i++ {
		allTicketsValid := r.isValid(i, tickets)
		if allTicketsValid {
			validIndexes = append(validIndexes, i)
		}
	}
	return validIndexes
}
