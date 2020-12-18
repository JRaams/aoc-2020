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
	a := evaluateExpressions(lines)
	fmt.Printf("Solution day 18 part a: %d\n", a)
}

func evaluateExpressions(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += parse(line)
	}
	return sum
}

//
// Taken from Arknave: https://github.com/arknave/advent-of-code-2020/blob/main/day18/day18.py
// Translated from python to Go
//
func parse(eqn string) int {
	// eqn: 1 + 2 * 3 + 4 * 5 + 6
	eqn = strings.ReplaceAll(eqn, "(", "( ")
	eqn = strings.ReplaceAll(eqn, ")", " )")

	tokens := funk.ReverseStrings(strings.Split(eqn, " "))
	// tokens: [1 + 2 * 3 + 4 * 5 + 6]
	stk := []string{}
	ops := []string{}

	for _, c := range tokens {
		token := string(c)
		if token == "(" {
			if len(ops) >= 1 {
				lastOp := ops[len(ops)-1]
				for lastOp != ")" {
					stk = append(stk, ops[len(ops)-1])
					ops = ops[:len(ops)-1]
					lastOp = ops[len(ops)-1]
				}
				ops = ops[:len(ops)-1]
			}
		} else if token == "*" || token == ")" || token == "+" {
			ops = append(ops, token)
		} else {
			stk = append(stk, token)
		}
	}
	// stk: [6 5 4 3 2 1]
	// ops: [+ * + * +]

	for len(ops) > 0 {
		stk = append(stk, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}
	// stk: [6 5 4 3 2 1 + * + * +]
	// ops: []

	curr := []int{}
	for _, val := range stk {
		if val == "+" {
			x := curr[len(curr)-1] + curr[len(curr)-2]
			curr = curr[:len(curr)-1]
			curr = curr[:len(curr)-1]
			curr = append(curr, x)
		} else if val == "*" {
			x := curr[len(curr)-1] * curr[len(curr)-2]
			curr = curr[:len(curr)-1]
			curr = curr[:len(curr)-1]
			curr = append(curr, x)
		} else {
			curr = append(curr, helpers.MustAtoi(val))
		}
	}

	return curr[0]
}
