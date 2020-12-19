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
	regex := getRegex(rules, 0)

	// Part 1
	a := solveA(regex, messages)
	fmt.Printf("Solution day 19 part a: %d\n", a)
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

func getRegex(rules map[int]rule, rootRuleNum int) string {
	r := ""
	rootRule := rules[rootRuleNum]

	// Reached a char rule, return it
	if rootRule.char != "" {
		return rootRule.char
	}

	//  Rule contains multiple subrules
	if len(rootRule.subRuleNumLists) == 2 {
		r += "("
		for _, subRuleNum := range rootRule.subRuleNumLists[0] {
			r += getRegex(rules, subRuleNum)
		}
		r += "|"
		for _, subRuleNum := range rootRule.subRuleNumLists[1] {
			r += getRegex(rules, subRuleNum)
		}
		r += ")"
		return r
	}

	if len(rootRule.subRuleNumLists) == 1 {
		for _, subRuleNum := range rootRule.subRuleNumLists[0] {
			r += getRegex(rules, subRuleNum)
		}
	}

	return r
}

func solveA(regex string, messages []string) int {
	a := 0
	regex = "^" + regex + "$"

	r, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}

	for _, message := range messages {
		matches := r.Match([]byte(message))
		if matches {
			a++
		}
	}
	return a
}
