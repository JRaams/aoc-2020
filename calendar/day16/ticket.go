package main

import (
	"strings"

	"github.com/jraams/aoc-2020/helpers"
)

type rule struct {
	name        string
	range1start int
	range1end   int
	range2start int
	range2end   int
}

type ticket struct {
	values []int
}

func (t *ticket) isValid(rules []rule) (bool, int) {
	for _, value := range t.values {
		valid := doesAtLeastOneRuleApply(value, rules)
		if !valid {
			return false, value
		}
	}
	return true, -1
}

func doesAtLeastOneRuleApply(value int, rules []rule) bool {
	for _, rule := range rules {
		if rule.range1start <= value && value <= rule.range1end || rule.range2start <= value && value <= rule.range2end {
			return true
		}
	}
	return false
}

func loadTicketData(lines []string) ([]rule, ticket, []ticket) {
	var rules []rule
	var ownTicket ticket
	var nearbyTickets []ticket

	readRules := true
	readOwnTicket := false
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line == "your ticket:" {
			readRules = false
			readOwnTicket = true
			continue
		}
		if line == "nearby tickets:" {
			readOwnTicket = false
			continue
		}

		if readRules {
			splitLine := strings.Split(line, ": ")
			ranges := strings.Split(splitLine[1], " or ")
			range1 := strings.Split(ranges[0], "-")
			range2 := strings.Split(ranges[1], "-")
			rule := rule{
				name:        splitLine[0],
				range1start: helpers.MustAtoi(range1[0]),
				range1end:   helpers.MustAtoi(range1[1]),
				range2start: helpers.MustAtoi(range2[0]),
				range2end:   helpers.MustAtoi(range2[1]),
			}
			rules = append(rules, rule)
		} else {
			strValues := strings.Split(line, ",")
			intValues := helpers.TranslateStringArrToIntArr(strValues)
			ticket := ticket{
				values: intValues,
			}

			if readOwnTicket {
				ownTicket = ticket
			} else {
				nearbyTickets = append(nearbyTickets, ticket)
			}
		}
	}

	return rules, ownTicket, nearbyTickets
}
