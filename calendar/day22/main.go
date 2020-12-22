package main

import (
	"container/list"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	d1, d2 := loadDecks(lines)

	// Part a
	a := solveA(d1, d2)
	fmt.Printf("Solution day 22 part a: %d\n", a)
}

func loadDecks(lines []string) (*list.List, *list.List) {
	d1, d2 := list.New(), list.New()

	loadingDeck1 := true
	for _, line := range lines {
		if strings.Contains(line, "Player") {
			continue
		}
		if len(line) == 0 {
			loadingDeck1 = false
			continue
		}
		intVal := helpers.MustAtoi(line)
		if loadingDeck1 {
			d1.PushBack(intVal)
		} else {
			d2.PushBack(intVal)
		}
	}

	return d1, d2
}

func solveA(d1 *list.List, d2 *list.List) int {
	// Play game
	for d1.Len() > 0 && d2.Len() > 0 {
		p1e := d1.Front()
		p2e := d2.Front()

		p1i := p1e.Value.(int)
		p2i := p2e.Value.(int)

		if p1i > p2i {
			d1.PushBack(p1i)
			d1.PushBack(p2i)
		} else {
			d2.PushBack(p2i)
			d2.PushBack(p1i)
		}

		d1.Remove(p1e)
		d2.Remove(p2e)
	}

	// Decide winner
	var winningList *list.List
	if d1.Len() == 0 {
		winningList = d2
	} else {
		winningList = d1
	}

	// Calculate score for winner
	score := 0
	for mult := winningList.Len(); mult > 0; {
		val := winningList.Front()
		score += mult * val.Value.(int)
		winningList.Remove(val)
		mult--
	}
	return score
}
