package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	rules, messages := load(lines)

	// Part 1
	a := solve(rules, messages, false)
	fmt.Printf("Solution day 19 part a: %d\n", a)

	// Part 2
	b := solve(rules, messages, true)
	fmt.Printf("Solution day 19 part b: %d\n", b)
}

type rule struct {
	char            string
	subRuleNumLists [][]int
}

func load(lines []string) (map[int]rule, []string) {
	rules := make(map[int]rule)
	var messages []string

	loadingRules := true
	for _, line := range lines {
		if line == "" {
			loadingRules = false
			continue
		}

		if loadingRules {
			splitLine := strings.Split(line, ": ")
			num, _ := strconv.Atoi(splitLine[0])
			ruleStr := splitLine[1]
			newRule := rule{}

			// Multiple lists of subrules
			if strings.Contains(ruleStr, "|") {
				subRuleSets := strings.Split(ruleStr, " | ")
				newRule.subRuleNumLists = [][]int{}
				for i := 0; i < len(subRuleSets); i++ {
					newRule.subRuleNumLists = append(newRule.subRuleNumLists, helpers.TranslateStringArrToIntArr(strings.Split(subRuleSets[i], " ")))
				}

				// Match a single character
			} else if strings.Contains(ruleStr, "\"") {
				char := strings.ReplaceAll(ruleStr, "\"", "")
				newRule.char = char

				// Single list of Subrules
			} else {
				subRuleNums := helpers.TranslateStringArrToIntArr(strings.Split(ruleStr, " "))
				newRule.subRuleNumLists = [][]int{
					subRuleNums,
				}
			}

			rules[num] = newRule
		} else {
			messages = append(messages, line)
		}
	}

	return rules, messages
}

func getRegex(rules map[int]rule, rootRuleNum int, useFixedRules bool) string {
	r := ""
	rootRule := rules[rootRuleNum]

	// Fix rules 8 & 11
	if useFixedRules {
		if rootRuleNum == 8 {
			return getRegex(rules, 42, useFixedRules) + "+"
		}
		if rootRuleNum == 11 {
			return getRegex(rules, 42, useFixedRules) + "{n}" + getRegex(rules, 31, useFixedRules) + "{n}"
		}
	}

	// Reached a char rule, return it
	if rootRule.char != "" {
		return rootRule.char
	}

	//  Rule contains two subrules
	if len(rootRule.subRuleNumLists) == 2 {
		r += "("
		for _, subRuleNum := range rootRule.subRuleNumLists[0] {
			r += getRegex(rules, subRuleNum, useFixedRules)
		}
		r += "|"
		for _, subRuleNum := range rootRule.subRuleNumLists[1] {
			r += getRegex(rules, subRuleNum, useFixedRules)
		}
		r += ")"
		return r
	}

	// Rule is a single list of subrules
	if len(rootRule.subRuleNumLists) == 1 {
		for _, subRuleNum := range rootRule.subRuleNumLists[0] {
			r += getRegex(rules, subRuleNum, useFixedRules)
		}
	}

	return r
}

func solve(rules map[int]rule, messages []string, useFixedRules bool) int {
	regex := getRegex(rules, 0, useFixedRules)
	regex = "^" + regex + "$"
	matchedMessageCount := 0

	maxN := 10
	if useFixedRules {
		maxN = 1
	}

	for _, message := range messages {
		for n := 1; n <= maxN; n++ {
			quantificationAmount := strings.ReplaceAll(regex, "n", strconv.Itoa(n))
			r, _ := regexp.Compile(quantificationAmount)

			matches := r.Match([]byte(message))
			if matches {
				matchedMessageCount++
				break
			}
		}
	}

	return matchedMessageCount
}
