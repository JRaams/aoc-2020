package main

import (
	"fmt"
	"path/filepath"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	rules, _, nearbyTickets := loadTicketData(lines)

	a := solveA(nearbyTickets, rules)
	fmt.Printf("Solution day 16 part a: %d\n", a)
}

func solveA(nearbyTickets []ticket, rules []rule) int {
	errorRateSum := 0
	for _, ticket := range nearbyTickets {
		ticketIsValid, errorRate := ticket.isValid(rules)
		if !ticketIsValid {
			errorRateSum += errorRate
		}
	}
	return errorRateSum
}
