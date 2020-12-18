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

	// Part 1
	a := evaluateExpressions(lines, false)
	fmt.Printf("Solution day 18 part a: %d\n", a)

	// Part 2
	b := evaluateExpressions(lines, true)
	fmt.Printf("Solution day 18 part b: %d\n", b)
}

func evaluateExpressions(lines []string, useAdvancedMath bool) int {
	sum := 0
	for _, line := range lines {
		sum += evaluate(line, useAdvancedMath)
	}
	return sum
}

//
// Taken from Arknave: https://github.com/arknave/advent-of-code-2020/blob/main/day18/day18.py
// Translated from python to Go
//
func evaluate(line string, useAdvancedMath bool) int {
	rep := strings.NewReplacer("(", "( ", ")", " )")
	tokens := funk.ReverseStrings(strings.Split(rep.Replace(line), " "))
	stk := []string{}
	ops := []string{}

	for _, token := range tokens {
		if token == "(" {
			if len(ops) >= 1 {
				lastOp := ops[len(ops)-1]
				for lastOp != ")" {
					stk = append(stk, helpers.PopStr(&ops))
					lastOp = ops[len(ops)-1]
				}
				ops = ops[:len(ops)-1]
			}
		} else if token == "*" {
			if useAdvancedMath {
				for len(ops) >= 1 && ops[len(ops)-1] == "+" {
					stk = append(stk, helpers.PopStr(&ops))
				}
			}
			ops = append(ops, token)
		} else if token == ")" || token == "+" {
			ops = append(ops, token)
		} else {
			stk = append(stk, token)
		}
	}

	for len(ops) > 0 {
		stk = append(stk, helpers.PopStr(&ops))
	}

	curr := []int{}
	for _, val := range stk {
		if val == "+" {
			x := curr[len(curr)-1] + curr[len(curr)-2]
			curr = curr[:len(curr)-2]
			curr = append(curr, x)
		} else if val == "*" {
			x := curr[len(curr)-1] * curr[len(curr)-2]
			curr = curr[:len(curr)-2]
			curr = append(curr, x)
		} else {
			curr = append(curr, helpers.MustAtoi(val))
		}
	}

	return curr[0]
}
