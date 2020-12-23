package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)

	// Part a
	vca, currentCup := load(lines[0], false)
	a := solveA(vca, currentCup)
	fmt.Printf("Solution day 23 part a: %s\n", a)

	// Part b
	vcb, currentCup := load(lines[0], true)
	b := solveB(vcb, currentCup)
	fmt.Printf("Solution day 23 part b: %d\n", b)
}

func solveA(vc valueCupMap, currentCup *cup) string {
	playCrabGame(vc, currentCup, 100)

	str := ""
	for cup := vc[1].next; cup.value != 1; cup = cup.next {
		str += strconv.Itoa(cup.value)
	}
	return str
}

func solveB(vc valueCupMap, currentCup *cup) int {
	playCrabGame(vc, currentCup, 10000000)

	cup1 := vc[1]
	b := cup1.next.value * cup1.next.next.value
	return b
}
