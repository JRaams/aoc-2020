package main

import (
	"container/list"
	"fmt"
	"path/filepath"
	"time"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	d1, d2 := loadDecks(lines)

	// Part a
	a := solveA(copyDeck(d1), copyDeck(d2))
	fmt.Printf("Solution day 22 part a: %d\n", a)

	// Part b
	b := solveB(copyDeck(d1), copyDeck(d2))
	fmt.Printf("Solution day 22 part b: %d\n", b)
}

func solveA(d1 *list.List, d2 *list.List) int {
	defer helpers.Measure(time.Now(), "")
	gameRoundDeckMap := map[int]*map[int][]string{}
	_, winningDeck := playSpaceCards(1, false, gameRoundDeckMap, d1, d2)
	score := calculateScore(winningDeck)
	return score
}

func solveB(d1 *list.List, d2 *list.List) int {
	defer helpers.Measure(time.Now(), "")
	gameRoundDeckMap := map[int]*map[int][]string{}
	_, winningDeck := playSpaceCards(1, true, gameRoundDeckMap, d1, d2)
	score := calculateScore(winningDeck)
	return score
}
